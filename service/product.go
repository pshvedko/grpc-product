package service

import (
	"net/http"
	"time"
)

type Browser interface {
	Do(req *http.Request) (*http.Response, error)
}

type Storage interface{}

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
