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
	var reply *product.ListReply
	reply, err = c.List(context.TODO(), &product.ListQuery{Page: &product.Page{
		Limit:  10,
		Offset: 20,
	}, Sort: []*product.Sort{{
		Order: false,
		By:    "price",
	}, {
		Order: true,
		By:    "name",
	}}})
	if err != nil {
		log.Fatal(err)
	}
	encoder := json.NewEncoder(os.Stdout)
	encoder.SetIndent("", "  ")
	err = encoder.Encode(reply)
	if err != nil {
		log.Fatal(err)
	}
	_, err = c.Fetch(context.TODO(), &product.FetchQuery{Url: "http://example.com/"})
	if err != nil {
		log.Fatal(err)
	}
}
