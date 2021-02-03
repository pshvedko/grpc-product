package service

import (
	"context"
	"github.com/pshvedko/grpc-product/storage"
)

type ListQuery interface {
	GetLimit() uint32
	GetOffset() uint32
	GetSortField() []string
}

type ListReply interface {
	storage.Cursor
}

func (s Service) List(_ context.Context, query ListQuery) (ListReply, error) {
	return s.Products().List(query.GetLimit(), query.GetOffset(), query.GetSortField()), nil
}
