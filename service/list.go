package service

import (
	"context"
	"encoding/json"
	"github.com/pshvedko/grpc-product/storage"
	"time"
)

type ListQuery interface {
	storage.Pager
	storage.Sorter
}

type ListReply interface {
	storage.Iterator
}

type iter struct {
	items []Product
	count int
	error
}

func (i iter) Err() error {
	return i.error
}

func (i iter) Close() error {
	return i.Err()
}

func (i iter) Done() bool {
	if i.error != nil || i.count == len(i.items) {
		return true
	}
	return false
}

func (i *iter) Next(o interface{}) bool {
	if i.Done() {
		return false
	}
	b, err := json.Marshal(i.items[i.count])
	if err != nil {
		i.error = err
		return false
	}
	err = json.Unmarshal(b, o)
	if err != nil {
		i.error = err
		return false
	}
	i.count++
	return false
}

func (s Service) List(ctx context.Context, query ListQuery) (ListReply, error) {
	_, _ = s.Find(ctx, query, query)
	return &iter{items: []Product{
		{
			Name:    "1",
			Price:   1,
			Changes: 10,
			Date:    time.Now(),
		}, {
			Name:    "2",
			Price:   2,
			Changes: 20,
			Date:    time.Now(),
		}, {
			Name:    "3",
			Price:   3,
			Changes: 30,
			Date:    time.Now(),
		},
	}}, nil
}
