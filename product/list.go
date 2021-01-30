package product

import (
	"context"
	"github.com/pshvedko/grpc-product/service"
	"log"
)

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
	log.Printf("LIST: %v", query)
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
