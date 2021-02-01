package cmd

import (
	"context"
	"fmt"
	"github.com/pshvedko/grpc-product/product"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"net"
	"net/url"
)

var fetchCmd = &cobra.Command{
	Use:   "fetch url",
	Args:  cobra.ExactArgs(1),
	Short: "Load external CSV file",
	Long:  "Fetches a CSV file with the 'PRODUCT NAME; PRICE' format from the specified URL and saves it to MongoDB",
	RunE:  runFetch,
}

func runFetch(_ *cobra.Command, args []string) (err error) {
	var u *url.URL
	u, err = url.Parse(args[0])
	if err != nil {
		return
	}
	if u.Scheme == "" || u.Host == "" {
		return fmt.Errorf("invalid url, please specify as 'http://host/file.csv'")
	}
	addr := net.TCPAddr{
		IP:   addrFlag,
		Port: portFlag,
	}
	var dial *grpc.ClientConn
	dial, err = grpc.Dial(addr.String(), grpc.WithInsecure())
	if err != nil {
		return
	}
	defer dial.Close()
	client := product.NewProductServiceClient(dial)
	var reply *product.FetchReply
	reply, err = client.Fetch(context.TODO(), &product.FetchQuery{Url: args[0]})
	if err != nil {
		return
	}
	fmt.Printf("Node %d, fetched %d, created %d, updated %d rows\n",
		reply.Node,
		reply.Fetched,
		reply.Created,
		reply.Updated)
	return
}

func init() {
	rootCmd.AddCommand(fetchCmd)
}
