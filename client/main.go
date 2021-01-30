package main

import (
	"context"
	"encoding/json"
	"github.com/pshvedko/grpc-product/product"
	"google.golang.org/grpc"
	"log"
	"os"
)

func main() {
	dial, err := grpc.Dial(":9000", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer dial.Close()
	c := product.NewProductServiceClient(dial)
	var reply *product.FetchReply
	reply, err = c.Fetch(context.TODO(), &product.FetchQuery{Url: "http://example.com/"})
	if err != nil {
		log.Fatal(err)
	}
	encoder := json.NewEncoder(os.Stdout)
	encoder.SetIndent("", "  ")
	err = encoder.Encode(reply)
	if err != nil {
		log.Fatal(err)
	}
}
