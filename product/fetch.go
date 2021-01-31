package product

import (
	"context"
	"log"
)

func (s API) Fetch(ctx context.Context, query *FetchQuery) (*FetchReply, error) {
	if s.Service == nil {
		return nil, ErrService
	}
	loaded, changed, added, err := s.Service.Fetch(ctx, query)
	if err != nil {
		return nil, err
	}
	log.Printf("fetch: %v %v %v %v", query, loaded, changed, added)
	return &FetchReply{Size: loaded}, nil
}
