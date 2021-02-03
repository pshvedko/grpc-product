package storage

import (
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
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
		s.acquire().Close()
	}
}

type Table interface {
	List(limit, offset uint32, sort []string) Cursor
}

type Products interface {
	Push(string, float64) (int, error)
	Table
}

type cursor struct {
	t *table
	i *mgo.Iter
}

func (c cursor) Next(v interface{}) bool {
	return c.i.Next(v)
}

func (c cursor) Close() error {
	defer c.t.release()
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

func (s *Storage) Products() Products {
	return products{table{s, s.acquire().DB("").C(defaultProductCollection)}}
}

func (s *Storage) acquire() *mgo.Session {
	return <-s.sessions
}

type A []interface{}

type table struct {
	s *Storage
	c *mgo.Collection
}

func (t table) release() {
	t.s.sessions <- t.c.Database.Session
}

func (t table) List(limit, offset uint32, sort []string) Cursor {
	return cursor{&t, t.c.Find(nil).Select(nil).Sort(sort...).Limit(int(limit)).Skip(int(offset)).Iter()}
}

type products struct {
	table
}

func (p products) Push(name string, price float64) (int, error) {
	defer p.release()
	info, err := p.c.Upsert(bson.M{"name": name},
		[]bson.M{{"$set": bson.M{
			"date": bson.M{
				"$cond": A{bson.M{"$ne": A{"$price", price}}, "$$NOW", "$date"}},
			"changes": bson.M{
				"$sum": A{bson.M{"$cond": A{bson.M{"$ne": A{"$price", price}}, 1, 0}}, "$changes"}},
			"price": price}}})
	if err != nil {
		return 0, err
	}
	return info.Updated + info.Matched, nil
}
