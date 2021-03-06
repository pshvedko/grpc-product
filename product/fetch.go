package product

import (
	"context"
	"log"
)

func (s Application) Fetch(ctx context.Context, query *FetchQuery) (*FetchReply, error) {
	if s.Service == nil {
		return nil, ErrService
	}
	loaded, changed, added, err := s.Service.Fetch(ctx, query)
	log.Printf("fetch: %v %v %v %v %v", query, loaded, changed, added, err)
	if err != nil {
		return nil, err
	}
	return &FetchReply{Fetched: loaded, Created: added, Updated: changed, Node: s.Id}, nil
}
