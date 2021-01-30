package service

import (
	"context"
	"encoding/json"
	"time"
)

//
//if query == nil {
//	return nil, ErrQuery
//}
//request, err := http.NewRequestWithContext(ctx, http.MethodGet, query.Url, nil)
//var response *http.Response
//response, err = s.Do(request)
//if err != nil {
//	return nil, err
//}
//defer response.Body.Close()
//
//r := csv.NewReader(response.Body)
//r.Comma = ';'
//r.Comment = '#'
//r.FieldsPerRecord = 2
//r.LazyQuotes = false
//r.TrimLeadingSpace = true
//
//var records []string
//records, err = r.Read()
//if err != nil {
//	return nil, err
//}

type FetchQuery interface {
	GetUrl() string
}

type ListQuery interface {
	GetPage() Page
	GetSort() []Sort
}

type Page interface {
	GetLimit() uint32
	GetOffset() uint32
}

type Sort interface {
	GetOrder() bool
	GetBy() string
}

type ListReply interface {
	Next(interface{}) bool
	Close() error
	Done() bool
	Err() error
}

type Service struct {
}

func (s Service) Fetch(ctx context.Context, query FetchQuery) (uint32, error) {
	return 0, nil
}

type Product struct {
	Name    string    `json:"name"`
	Price   string    `json:"price"`
	Changes uint32    `json:"changes"`
	Date    time.Time `json:"date"`
}

type iter struct {
	items []Product
	count int
	error
}

func (i iter) Err() error {
	return i.error
}

func (i iter) Close() error {
	return i.Err()
}

func (i iter) Done() bool {
	if i.error != nil || i.count == len(i.items) {
		return true
	}
	return false
}

func (i *iter) Next(o interface{}) bool {
	if i.Done() {
		return false
	}
	b, err := json.Marshal(i.items[i.count])
	if err != nil {
		i.error = err
		return false
	}
	err = json.Unmarshal(b, o)
	if err != nil {
		i.error = err
		return false
	}
	i.count++
	return false
}

func (s Service) List(ctx context.Context, query ListQuery) (ListReply, error) {
	return &iter{items: []Product{
		{
			Name:    "1",
			Price:   "1",
			Changes: 10,
			Date:    time.Now(),
		}, {
			Name:    "2",
			Price:   "2",
			Changes: 20,
			Date:    time.Now(),
		}, {
			Name:    "3",
			Price:   "3",
			Changes: 30,
			Date:    time.Now(),
		},
	}}, nil
}
