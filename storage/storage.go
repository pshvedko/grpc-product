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
	defaultPoolSize = 16
	defaultDataBase = "foo"
)

func (s *Storage) Start() (err error) {
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
	_, err = migrate(session.DB(""))
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

type Table struct {
	s *Storage
	c *mgo.Collection
}

func (t Table) Release() {
	t.s.sessions <- t.c.Database.Session
}

func (t Table) Upsert(selector, update interface{}) (*mgo.ChangeInfo, error) {
	defer t.Release()
	return t.c.Upsert(selector, update)
}

func (t Table) Cursor(selector, field interface{}, limit, offset uint32, sort []string) *mgo.Iter {
	defer t.Release()
	return t.c.Find(selector).Select(field).Sort(sort...).Limit(int(limit)).Skip(int(offset)).Iter()
}

func (s *Storage) Products() Table {
	return s.Table("products")
}

func (s *Storage) Table(name string) Table {
	return Table{s, s.Acquire().DB("").C(name)}
}

func (s *Storage) Acquire() *mgo.Session {
	return <-s.sessions
}
