package cmd

import (
	"fmt"
	"net"

	"github.com/pshvedko/grpc-product/product"
	"github.com/pshvedko/grpc-product/service"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
)

var serveCmd = &cobra.Command{
	Use:   "serve",
	Args:  cobra.ExactArgs(0),
	Short: "Operate in service mode",
	Long:  `Starts operating in service mode with Fetch(url) and List(order, sort) methods`,
	RunE:  runServe,
}

func runServe(cmd *cobra.Command, args []string) (err error) {
	var listener net.Listener
	listener, err = net.Listen("tcp", fmt.Sprintf(":%v", portFlag))
	if err != nil {
		return
	}
	defer listener.Close()
	api := product.Server{
		Service: &service.Service{},
	}
	server := grpc.NewServer()
	product.RegisterProductServiceServer(server, api)
	err = server.Serve(listener)
	return
}

func init() {
	rootCmd.AddCommand(serveCmd)
}
