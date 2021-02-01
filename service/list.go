package service

import (
	"context"
)

type ListQuery interface {
	GetLimit() uint32
	GetOffset() uint32
	GetSortField() []string
}

type ListReply interface {
	Next(interface{}) bool
	Close() error
	Done() bool
	Err() error
}

func (s Service) List(_ context.Context, query ListQuery) (ListReply, error) {
	return s.Products().Cursor(nil, nil, query.GetLimit(), query.GetOffset(), query.GetSortField()), nil
}
