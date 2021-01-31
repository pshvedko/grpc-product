package product

import (
	"context"
	"log"
)

func (s Server) List(ctx context.Context, query *ListQuery) (*ListReply, error) {
	log.Printf("list: %v", query)
	if s.Service == nil {
		return nil, ErrService
	}
	rows, err := s.Service.List(ctx, query)
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
