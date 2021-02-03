package service

import (
	"context"
	"encoding/csv"
	"errors"
	"io"
	"net/http"
	"strconv"
)

type FetchQuery interface {
	GetUrl() string
}

var (
	ErrQuery  = errors.New("query is nil")
	ErrStatus = errors.New("can't load file")
)

func (s Service) Fetch(ctx context.Context, query FetchQuery) (loaded, changed, added uint32, err error) {
	if query == nil {
		return 0, 0, 0, ErrQuery
	}
	var request *http.Request
	request, err = http.NewRequestWithContext(ctx, http.MethodGet, query.GetUrl(), nil)
	var response *http.Response
	response, err = s.Do(request)
	if err != nil {
		return 0, 0, 0, err
	}
	defer response.Body.Close()
	if response.StatusCode != 200 {
		return 0, 0, 0, ErrStatus
	}
	r := csv.NewReader(response.Body)
	r.Comma = ';'
	r.Comment = '#'
	r.FieldsPerRecord = 2
	r.TrimLeadingSpace = true
	r.ReuseRecord = true
	r.LazyQuotes = true
	var push int
	var price float64
	first := true
	for {
		var record []string
		record, err = r.Read()
		if err != nil {
			if err == io.EOF {
				err = nil
			}
			break
		}
		if first {
			first = false
			continue
		}
		price, err = strconv.ParseFloat(record[1], 64)
		if err != nil {
			break
		}
		push, err = s.Products().Push(record[0], price)
		if err != nil {
			break
		}
		switch push {
		case 0:
			added++
		case 2:
			changed++
		}
		loaded++
	}
	return
}
