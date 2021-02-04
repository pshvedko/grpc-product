package product

import (
	"context"
	"log"
)

func (a Application) Fetch(ctx context.Context, query *FetchQuery) (*FetchReply, error) {
	if a.Service == nil {
		return nil, ErrService
	}
	loaded, changed, added, err := a.Service.Fetch(ctx, query)
	log.Printf("fetch: %v %v %v %v %v", query, loaded, changed, added, err)
	if err != nil {
		return nil, err
	}
	return &FetchReply{Fetched: loaded, Created: added, Updated: changed, Node: a.Id}, nil
}
