package product

//go:generate protoc -I . --go_out=plugins=grpc:. product.proto

import (
	"context"
	"encoding/json"
	"github.com/pshvedko/grpc-product/service"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
	"time"
)

type Service interface {
	Fetch(context.Context, service.FetchQuery) (uint32, error)
	List(context.Context, service.ListQuery) (service.ListReply, error)
	service.Starter
}

type API struct {
	Service
}

var (
	ErrService = status.Error(codes.Internal, "service is not defined")
)

func (x *ListQuery) GetLimit() uint32 {
	return x.Page.Limit
}

func (x *ListQuery) GetOffset() uint32 {
	return x.Page.Offset
}

func (x *ListQuery) ForSort(f func(string, bool)) {
	for _, v := range x.Sort {
		f(v.By, v.Order)
	}
}

type product struct {
	x    interface{}
	Date time.Time `json:"date"`
}

func (x *Product) UnmarshalJSON(data []byte) (err error) {
	v := product{x: x}
	err = json.Unmarshal(data, &v)
	if err == nil {
		x.Date = timestamppb.New(v.Date)
	}
	return
}

func (x *Product) MarshalJSON() ([]byte, error) {
	return json.Marshal(product{
		x:    x,
		Date: x.Date.AsTime(),
	})
}
