package product

//go:generate protoc -I . --go-grpc_out=. product.proto
//go:generate protoc -I . --go_out=. product.proto

import (
	"context"
	"encoding/json"
	"github.com/globalsign/mgo/bson"
	"github.com/pshvedko/grpc-product/service"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type Service interface {
	Fetch(context.Context, service.FetchQuery) (uint32, uint32, uint32, error)
	List(context.Context, service.ListQuery) (service.ListReply, error)
	service.Starter
}

type Application struct {
	Service
	Id uint32
}

func (s Application) Start() error {
	return s.Service.Start(s.Id)
}

func (s Application) mustEmbedUnimplementedProductServiceServer() {}

var (
	ErrService = status.Error(codes.Internal, "service is not defined")
)

func (x *ListQuery) GetLimit() uint32 {
	return x.Page.Limit
}

func (x *ListQuery) GetOffset() uint32 {
	return x.Page.Offset
}

func (x *ListQuery) GetSortField() (fields []string) {
	for _, v := range x.Sort {
		var minus string
		if v.Order {
			minus = "-"
		}
		fields = append(fields, minus+v.By)
	}
	return
}

func NewSort(field string) *Sort {
	var minus bool
	switch field[0] {
	case '-':
		minus = true
		fallthrough
	case '+':
		field = field[1:]
	}
	return &Sort{
		Order: minus,
		By:    field,
	}
}

func (x *Product) set(v service.Product) {
	x.Name = v.Name
	x.Price = v.Price
	x.Changes = v.Changes
	x.Date = timestamppb.New(v.Date)
}

func (x *Product) get() service.Product {
	return service.Product{
		Name:    x.Name,
		Price:   x.Price,
		Changes: x.Changes,
		Date:    x.Date.AsTime(),
	}
}

func (x *Product) UnmarshalJSON(data []byte) (err error) {
	var v service.Product
	err = json.Unmarshal(data, &v)
	if err == nil {
		x.set(v)
	}
	return
}

func (x *Product) MarshalJSON() ([]byte, error) {
	return json.Marshal(x.get())
}

func (x *Product) SetBSON(raw bson.Raw) (err error) {
	var v service.Product
	err = raw.Unmarshal(&v)
	if err == nil {
		x.set(v)
	}
	return
}

func (x *Product) GetBSON() (interface{}, error) {
	return x.get(), nil
}
