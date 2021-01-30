package product

//go:generate protoc -I . --go_out=plugins=grpc:. product.proto

import (
	"context"
	"encoding/json"
	"github.com/pshvedko/grpc-product/service"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
	"log"
)

type Service interface {
	Fetch(context.Context, service.FetchQuery) (uint32, error)
	List(context.Context, service.ListQuery) (service.ListReply, error)
}

type Server struct {
	Service
}

var (
	ErrService = status.Error(codes.Internal, "service is not defined")
)

func (s Server) Fetch(ctx context.Context, query *FetchQuery) (*FetchReply, error) {
	log.Printf("%v", query)
	if s.Service == nil {
		return nil, ErrService
	}
	size, err := s.Service.Fetch(ctx, query)
	if err != nil {
		return nil, err
	}
	return &FetchReply{Size: size}, nil
}

type ListQueryService struct {
	*ListQuery
}

func (q ListQueryService) GetPage() service.Page {
	return q.Page
}

func (q ListQueryService) GetSort() []service.Sort {
	var p []service.Sort
	for _, v := range q.Sort {
		p = append(p, v)
	}
	return p
}

func (s Server) List(ctx context.Context, query *ListQuery) (*ListReply, error) {
	log.Printf("%v", query)
	if s.Service == nil {
		return nil, ErrService
	}
	rows, err := s.Service.List(ctx, ListQueryService{query})
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var products []*Product
	for !rows.Done() {
		var row Product
		if !rows.Next(&row) {
			err = rows.Err()
			if err != nil {
				return nil, err
			}
		}
		products = append(products, &row)
	}
	return &ListReply{Products: products}, nil
}

func (x *Product) UnmarshalJSON(data []byte) (err error) {
	var v service.Product
	err = json.Unmarshal(data, &v)
	if err != nil {
		return
	}
	x.Name, x.Price, x.Changes, x.Date = v.Name, v.Price, v.Changes, timestamppb.New(v.Date)
	return
}

func (x *Product) MarshalJSON() ([]byte, error) {
	return json.Marshal(service.Product{
		Name:    x.Name,
		Price:   x.Price,
		Changes: x.Changes,
		Date:    x.Date.AsTime(),
	})
}
