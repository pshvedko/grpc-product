package service

import (
	"context"
	"net/http"
	"time"

	"github.com/pshvedko/grpc-product/storage"
)

type Browser interface {
	Do(*http.Request) (*http.Response, error)
}

type Storage interface {
	Add(context.Context, storage.Product) error
	Find(context.Context, storage.Pager, storage.Sorter) (storage.Iterator, error)
}

type Service struct {
	Browser
	Storage
}

type Product struct {
	Name    string    `json:"name"`
	Price   float64   `json:"price"`
	Changes uint32    `json:"changes"`
	Date    time.Time `json:"date"`
}
