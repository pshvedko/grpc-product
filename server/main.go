package main

import (
	"log"
	"net"

	"github.com/pshvedko/grpc-product/product"
	"google.golang.org/grpc"
)

func main() {
	listener, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()
	server := grpc.NewServer()
	product.RegisterProductServiceServer(server, product.Server{})
	if err := server.Serve(listener); err != nil {
		log.Fatal(err)
	}
}
