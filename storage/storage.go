package storage

import (
	"context"
	"github.com/globalsign/mgo"
)

type Product interface{}

type Pager interface {
	GetLimit() uint32
	GetOffset() uint32
}

type Sorter interface {
	ForSort(func(string, bool))
}

type Iterator interface {
	Next(interface{}) bool
	Close() error
	Done() bool
	Err() error
}

type Mongo mgo.DialInfo

func (m *Mongo) Info() *mgo.DialInfo {
	return (*mgo.DialInfo)(m)
}

type Storage struct {
	*Mongo
	db *mgo.Database
}

func (m *Storage) Start() error {
	ses, err := mgo.DialWithInfo(m.Info())
	if err != nil {
		return err
	}
	m.db = ses.DB(m.Database)
	_, err = migrate(m.db)
	if err != nil {
		m.Stop()
	}
	for i := 0; i < m.PoolLimit; i++ {
	}
	return err
}

func (m *Storage) Stop() {
	m.db.Session.Close()
}

func (m *Storage) Add(_ context.Context, product Product) error {
	return m.db.C("products").Insert(product)
}

func (Storage) Find(context.Context, Pager, Sorter) (Iterator, error) {
	return nil, nil
}
