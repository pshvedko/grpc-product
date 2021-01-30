package product

//go:generate protoc -I . --go_out=plugins=grpc:. product.proto

import (
	"context"
	"encoding/json"
	"github.com/pshvedko/grpc-product/service"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
	"strconv"
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

func (x *Product) UnmarshalJSON(data []byte) (err error) {
	var v service.Product
	price := strconv.FormatFloat(v.Price, 'f', -1, 64)
	err = json.Unmarshal(data, &v)
	if err != nil {
		return
	}
	x.Name, x.Price, x.Changes, x.Date = v.Name, price, v.Changes, timestamppb.New(v.Date)
	return
}

func (x *Product) MarshalJSON() ([]byte, error) {
	price, err := strconv.ParseFloat(x.Price, 64)
	if err != nil {
		return nil, err
	}
	return json.Marshal(service.Product{
		Name:    x.Name,
		Price:   price,
		Changes: x.Changes,
		Date:    x.Date.AsTime(),
	})
}
