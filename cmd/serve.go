package cmd

import (
	"fmt"
	"math/rand"
	"net"
	"net/http"
	"time"

	"github.com/pshvedko/grpc-product/product"
	"github.com/pshvedko/grpc-product/service"
	"github.com/pshvedko/grpc-product/storage"
	"github.com/spf13/cobra"

	"google.golang.org/grpc"
)

var serveCmd = &cobra.Command{
	Use:   "serve",
	Args:  cobra.ExactArgs(0),
	Short: "Operate in service mode",
	Long:  "Starts operating in service mode with Fetch(url) and List(order, sort) methods",
	RunE:  runServe,
}

func runServe(*cobra.Command, []string) (err error) {
	var listener net.Listener
	listener, err = net.Listen("tcp", fmt.Sprintf(":%v", portFlag))
	if err != nil {
		return
	}
	defer listener.Close()
	api := &product.Application{
		Service: &service.Service{
			Browser: &http.Client{},
			Storage: &storage.Storage{
				Mongo: &storage.Mongo{
					Addrs:     []string{"192.168.0.244:27017"},
					Database:  "foo",
					Username:  "",
					Password:  "",
					PoolLimit: 10,
				},
			},
		},
		Id: nodeFlag,
	}
	err = api.Start()
	if err != nil {
		return
	}
	defer api.Stop()
	server := grpc.NewServer()
	product.RegisterProductServiceServer(server, api)
	err = server.Serve(listener)
	return
}

var nodeFlag uint32

func init() {
	n := time.Now().UnixNano()
	r := rand.NewSource(n)
	serveCmd.Flags().Uint32VarP(&nodeFlag, "node", "n", rand.New(r).Uint32() % 10, "node id")
	rootCmd.AddCommand(serveCmd)
}
