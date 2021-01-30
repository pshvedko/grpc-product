package product

//go:generate protoc -I . --go_out=plugins=grpc:. product.proto

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
	"log"
)

type Server struct{}

func (s Server) Fetch(_ context.Context, query *FetchQuery) (*FetchReply, error) {
	log.Printf("%v", query)
	return nil, status.Error(codes.Unimplemented, "not implemented yet")
}

func (s Server) List(_ context.Context, query *ListQuery) (*ListReply, error) {
	log.Printf("%v", query)
	return &ListReply{Products: []*Product{
		{
			Name:    "1",
			Price:   "1",
			Changes: 10,
			Date:    timestamppb.Now(),
		},
		{
			Name:    "2",
			Price:   "2",
			Changes: 20,
			Date:    timestamppb.Now(),
		},
	}}, nil
}
