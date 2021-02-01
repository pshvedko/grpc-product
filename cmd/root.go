package cmd

import (
	"fmt"
	"github.com/pshvedko/grpc-product/product"
	"github.com/pshvedko/grpc-product/service"
	"github.com/pshvedko/grpc-product/storage"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"net"
	"net/http"
	"os"
	"path/filepath"
)

import _ "github.com/pshvedko/grpc-product/migrate"

var portFlag int
var addrFlag net.IP

var rootCmd = &cobra.Command{
	Use:           filepath.Base(os.Args[0]),
	SilenceErrors: true,
	RunE:          runServe,
}

func runServe(cmd *cobra.Command, _ []string) (err error) {
	flag := cmd.Flag("node")
	if flag != nil && flag.Changed && !serverFlag {
		return fmt.Errorf("malformed usage flag: -n %s", flag.Usage)
	} else if !serverFlag {
		return cmd.Usage()
	}
	addr := net.TCPAddr{
		IP:   addrFlag,
		Port: portFlag,
	}
	var listener net.Listener
	listener, err = net.ListenTCP("tcp", &addr)
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

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

var serverFlag bool
var nodeFlag uint32

func init() {
	rootCmd.Flags().BoolVarP(&serverFlag, "", "s", false, "run in service mode")
	rootCmd.Flags().Uint32VarP(&nodeFlag, "node", "n", 0, "node id used with -s")
	rootCmd.PersistentFlags().IntVarP(&portFlag, "port", "p", 9000, "port to listen")
	rootCmd.PersistentFlags().IPVarP(&addrFlag, "addr", "a", net.IP{0, 0, 0, 0}, "address to bind")
}
