package service

import (
	"context"
	"encoding/csv"
	"errors"
	"github.com/globalsign/mgo/bson"
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

func (s Service) Fetch(ctx context.Context, query FetchQuery) (uint32, error) {
	if query == nil {
		return 0, ErrQuery
	}
	request, err := http.NewRequestWithContext(ctx, http.MethodGet, query.GetUrl(), nil)
	var response *http.Response
	response, err = s.Do(request)
	if err != nil {
		return 0, err
	}
	defer response.Body.Close()
	if response.StatusCode != 200 {
		return 0, ErrStatus
	}
	r := csv.NewReader(response.Body)
	r.Comma = ';'
	r.Comment = '#'
	r.FieldsPerRecord = 2
	r.TrimLeadingSpace = true
	r.ReuseRecord = true
	r.LazyQuotes = true
	var price float64
	var count uint32
	var products []Product
	for {
		var record []string
		record, err = r.Read()
		if err != nil {
			if err == io.EOF {
				break
			}
			return 0, err
		}
		price, err = strconv.ParseFloat(record[1], 64)
		if err != nil {
			return 0, err
		}
		products = append(products, Product{
			Name:  record[0],
			Price: price,
		})
		count++
	}
	type A []interface{}
	for _, product := range products {
		_, err = s.Products().Upsert(bson.M{"name": product.Name},
			[]bson.M{{"$set": bson.M{
				"price": product.Price,
				"date": bson.M{
					"$cond": A{bson.M{"$ne": A{"$price", product.Price}}, "$$NOW", "$date"}},
				"changes": bson.M{
					"$sum": A{bson.M{"$cond": A{bson.M{"$ne": A{"$price", product.Price}}, 1, 0}}, "$changes"}}}}})
		if err != nil {
			return 0, err
		}
	}
	return count, nil
}
