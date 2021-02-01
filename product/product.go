package product

//go:generate protoc -I . --go_out=plugins=grpc:. product.proto

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

type API struct {
	Service
}

var (
	ErrService = status.Error(codes.Internal, "service is not defined")
)

func (m *ListQuery) GetLimit() uint32 {
	return m.Page.Limit
}

func (m *ListQuery) GetOffset() uint32 {
	return m.Page.Offset
}

func (m *ListQuery) GetSortField() (fields []string) {
	for _, v := range m.Sort {
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

func (m *Product) set(v service.Product) {
	m.Name = v.Name
	m.Price = v.Price
	m.Changes = v.Changes
	m.Date = timestamppb.New(v.Date)
}

func (m *Product) get() service.Product {
	return service.Product{
		Name:    m.Name,
		Price:   m.Price,
		Changes: m.Changes,
		Date:    m.Date.AsTime(),
	}
}

func (m *Product) UnmarshalJSON(data []byte) (err error) {
	var v service.Product
	err = json.Unmarshal(data, &v)
	if err == nil {
		m.set(v)
	}
	return
}

func (m *Product) MarshalJSON() ([]byte, error) {
	return json.Marshal(m.get())
}

func (m *Product) SetBSON(raw bson.Raw) (err error) {
	var v service.Product
	err = raw.Unmarshal(&v)
	if err == nil {
		m.set(v)
	}
	return
}

func (m *Product) GetBSON() (interface{}, error) {
	return m.get(), nil
}
