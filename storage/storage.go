package storage

import (
	"github.com/globalsign/mgo"
)

type Mongo mgo.DialInfo

func (m *Mongo) Info() *mgo.DialInfo {
	return (*mgo.DialInfo)(m)
}

type Storage struct {
	*Mongo
	sessions chan *mgo.Session
}

const (
	defaultPoolSize          = 16
	defaultDataBase          = "foo"
	defaultProductCollection = "products"
)

func (s *Storage) Start(id uint32) (err error) {
	if s.PoolLimit == 0 {
		s.PoolLimit = defaultPoolSize
	}
	if s.Database == "" {
		s.Database = defaultDataBase
	}
	var session *mgo.Session
	session, err = mgo.DialWithInfo(s.Info())
	if err != nil {
		return err
	}
	_, err = migrate(session.DB(""), id)
	if err != nil {
		session.Close()
		return err
	}
	s.sessions = make(chan *mgo.Session, s.PoolLimit)
	for i := 0; i < s.PoolLimit-1; i++ {
		s.sessions <- session.Copy()
	}
	s.sessions <- session
	return err
}

func (s *Storage) Stop() {
	for i := 0; i < s.PoolLimit; i++ {
		s.Acquire().Close()
	}
}

type Table interface {
	Release()
	Upsert(selector, update interface{}) (*mgo.ChangeInfo, error)
	Cursor(selector, field interface{}, limit, offset uint32, sort []string) Cursor
}

type table struct {
	s *Storage
	c *mgo.Collection
}

func (t table) Release() {
	t.s.sessions <- t.c.Database.Session
}

func (t table) Upsert(selector, update interface{}) (*mgo.ChangeInfo, error) {
	defer t.Release()
	return t.c.Upsert(selector, update)
}

type cursor struct {
	t *table
	i *mgo.Iter
}

func (c cursor) Next(v interface{}) bool {
	return c.i.Next(v)
}

func (c cursor) Close() error {
	defer c.t.Release()
	return c.i.Close()
}

func (c cursor) Done() bool {
	return c.i.Done()
}

func (c cursor) Err() error {
	return c.i.Err()
}

type Cursor interface {
	Next(interface{}) bool
	Close() error
	Done() bool
	Err() error
}

func (t table) Cursor(selector, field interface{}, limit, offset uint32, sort []string) Cursor {
	return cursor{&t, t.c.Find(selector).Select(field).Sort(sort...).Limit(int(limit)).Skip(int(offset)).Iter()}
}

func (s *Storage) Products() Table {
	return s.Table(defaultProductCollection)
}

func (s *Storage) Table(name string) Table {
	return table{s, s.Acquire().DB("").C(name)}
}

func (s *Storage) Acquire() *mgo.Session {
	return <-s.sessions
}
