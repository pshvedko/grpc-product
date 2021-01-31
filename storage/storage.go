package storage

import "context"

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

type Mongo struct {
}

func (Mongo) Add(context.Context, Product) error {
	return nil
}

func (Mongo) Find(context.Context, Pager, Sorter) (Iterator, error) {
	return nil, nil
}
