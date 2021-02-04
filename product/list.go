package product

import (
	"context"
	"log"
)

func (a Application) List(ctx context.Context, query *ListQuery) (*ListReply, error) {
	if a.Service == nil {
		return nil, ErrService
	}
	rows, err := a.Service.List(ctx, query)
	log.Printf("list: %v %v", query, err)
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
	return &ListReply{Products: products, Node: a.Id}, nil
}
