package service

import (
	"context"
	"encoding/csv"
	"errors"
	"github.com/globalsign/mgo"
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

func (s Service) Fetch(ctx context.Context, query FetchQuery) (uint32, uint32, uint32, error) {
	if query == nil {
		return 0, 0, 0, ErrQuery
	}
	request, err := http.NewRequestWithContext(ctx, http.MethodGet, query.GetUrl(), nil)
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
	var price float64
	var changed uint32
	var loaded uint32
	var added uint32
	var info *mgo.ChangeInfo
	type A []interface{}
	for {
		var record []string
		record, err = r.Read()
		if err != nil {
			if err == io.EOF {
				err = nil
			}
			break
		}
		price, err = strconv.ParseFloat(record[1], 64)
		if err != nil {
			break
		}
		info, err = s.Products().Upsert(bson.M{"name": record[0]},
			[]bson.M{{"$set": bson.M{
				"date": bson.M{
					"$cond": A{bson.M{"$ne": A{"$price", price}}, "$$NOW", "$date"}},
				"changes": bson.M{
					"$sum": A{bson.M{"$cond": A{bson.M{"$ne": A{"$price", price}}, 1, 0}}, "$changes"}},
				"price": price}}})
		if err != nil {
			break
		}
		switch info.Matched + info.Updated {
		case 0:
			added++
		case 2:
			changed++
		}
		loaded++
	}
	return loaded, changed, added, err
}
