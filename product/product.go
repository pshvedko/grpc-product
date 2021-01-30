package product

//go:generate protoc -I . --go_out=plugins=grpc:. product.proto

import (
	"context"
	"errors"
	"log"
)

type Server struct{}

func (s Server) Fetch(_ context.Context, query *FetchQuery) (*FetchReply, error) {
	log.Print(query.Url)
	return &FetchReply{Size: 3}, nil
}

func (s Server) List(_ context.Context, query *ListQuery) (*ListReply, error) {
	log.Print(query.Page, query.Sort)
	return nil, errors.New("not found")
}
