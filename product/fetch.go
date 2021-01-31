package product

import (
	"context"
	"log"
)

func (s Server) Fetch(ctx context.Context, query *FetchQuery) (*FetchReply, error) {
	log.Printf("fetch: %v", query)
	if s.Service == nil {
		return nil, ErrService
	}
	size, err := s.Service.Fetch(ctx, query)
	if err != nil {
		return nil, err
	}
	return &FetchReply{Size: size}, nil
}
