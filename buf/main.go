package main

import (
	"context"
	"fmt"
	"net/http"

	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"

	"google.golang.org/protobuf/types/known/timestamppb"

	"connectrpc.com/connect"

	"github.com/pshvedko/grpc-product/gen/product"
	"github.com/pshvedko/grpc-product/gen/product/productconnect"
)

type ProductServer struct{}

func (p ProductServer) Fetch(ctx context.Context, c *connect.Request[product.FetchQuery]) (*connect.Response[product.FetchReply], error) {
	//TODO implement me
	panic("implement me")
}

func (p ProductServer) List(ctx context.Context, c *connect.Request[product.ListQuery]) (*connect.Response[product.ListReply], error) {
	res := connect.NewResponse(&product.ListReply{
		Products: func(o, l uint32) []*product.Product {
			var pp []*product.Product
			for l > 0 {
				pp = append(pp, &product.Product{
					Name:    fmt.Sprintf("product_%d", o),
					Price:   10,
					Changes: l,
					Date:    timestamppb.Now(),
				})
				l--
				o++
			}
			return pp
		}(c.Msg.Page.Offset, c.Msg.Page.Limit),
		Node: 1,
	})
	res.Header().Set("Product-Version", "v1")
	return res, nil
}

func (p ProductServer) ListStream(ctx context.Context, c *connect.BidiStream[product.ListQuery, product.Product]) error {
	//TODO implement me
	panic("implement me")
}

func main() {
	path, handler := productconnect.NewProductServiceHandler(&ProductServer{})
	mux := http.NewServeMux()
	mux.Handle(path, handler)
	http.ListenAndServe(":9000", h2c.NewHandler(mux, &http2.Server{}))
}
