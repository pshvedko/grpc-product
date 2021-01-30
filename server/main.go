package main

import (
	"github.com/pshvedko/grpc-product/product"
	"github.com/pshvedko/grpc-product/service"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()
	api := product.Server{
		Service: &service.Service{},
	}
	server := grpc.NewServer()
	product.RegisterProductServiceServer(server, api)
	if err := server.Serve(listener); err != nil {
		log.Fatal(err)
	}
}
