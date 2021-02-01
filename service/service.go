package service

import (
	"github.com/pshvedko/grpc-product/storage"
	"net/http"
	"time"
)

type Browser interface {
	Do(*http.Request) (*http.Response, error)
}

type Starter interface {
	Start(uint32) error
	Stop()
}

type Storage interface {
	Products() storage.Table
	Starter
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
