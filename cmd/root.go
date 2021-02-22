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
	"os/signal"
	"path/filepath"
	"syscall"
)

import _ "github.com/pshvedko/grpc-product/migrate"

var rootCmd = &cobra.Command{
	Use:           filepath.Base(os.Args[0]),
	SilenceErrors: true,
	RunE:          runServe,
}

func runServe(cmd *cobra.Command, _ []string) (err error) {
	if !serverFlag {
		return cmd.Usage()
	}
	var listener net.Listener
	listener, err = net.ListenTCP("tcp", &addressFlag)
	if err != nil {
		return
	}
	defer listener.Close()
	api := &product.Application{
		Service: &service.Service{
			Browser: &http.Client{},
			Storage: &storage.Storage{
				Mongo: &storage.Mongo{
					Addrs:     mongo.addressesFlag,
					Database:  mongo.nameFlag,
					Username:  mongo.userFlag,
					Password:  mongo.passFlag,
					PoolLimit: mongo.poolFlag,
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
	e := make(chan error)
	go func() {
		e <- server.Serve(listener)
	}()
	c := make(chan os.Signal)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
	defer close(c)
	defer close(e)
	for {
		select {
		case <-c:
			server.GracefulStop()
			signal.Stop(c)
		case err = <-e:
			return
		}
	}
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

var (
	addressFlag net.TCPAddr
	serverFlag  bool
	nodeFlag    uint32
	mongo       struct {
		addressesFlag []string
		nameFlag      string
		userFlag      string
		passFlag      string
		poolFlag      int
	}
)

func init() {
	rootCmd.Flags().BoolVarP(&serverFlag, "serve", "s", false, "run in service mode")
	rootCmd.Flags().Uint32VarP(&nodeFlag, "node", "n", 0, "node id")
	rootCmd.Flags().StringArrayVarP(&mongo.addressesFlag, "mongo", "m", []string{"mongo:27017"}, "mongo addresses")
	rootCmd.Flags().StringVarP(&mongo.nameFlag, "mongo-db-name", "N", "foo", "mongo db name")
	rootCmd.Flags().StringVarP(&mongo.userFlag, "mongo-user", "U", "", "mongo user")
	rootCmd.Flags().StringVarP(&mongo.passFlag, "mongo-pass", "P", "", "mongo password")
	rootCmd.Flags().IntVarP(&mongo.poolFlag, "mongo-pool", "S", 10, "mongo pool size")
	rootCmd.PersistentFlags().IntVarP(&addressFlag.Port, "port", "p", 9000, "port to listen")
	rootCmd.PersistentFlags().IPVarP(&addressFlag.IP, "addr", "a", net.IP{0, 0, 0, 0}, "address to bind")
}
