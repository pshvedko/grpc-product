package cmd

import (
	"context"
	"fmt"
	"github.com/pshvedko/grpc-product/product"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"net/url"
)

var fetchCmd = &cobra.Command{
	Use:   "fetch url",
	Args:  cobra.ExactArgs(1),
	Short: "Load external CSV file",
	Long: `The command loads a CSV file with the 'PRODUCT NAME; PRICE' format
from the specified URL and saves it to MongoDB`,
	RunE: runFetch,
}

func runFetch(cmd *cobra.Command, args []string) (err error) {
	_, err = url.Parse(args[0])
	if err != nil {
		return
	}
	var dial *grpc.ClientConn
	dial, err = grpc.Dial(fmt.Sprintf(":%v", portFlag), grpc.WithInsecure())
	if err != nil {
		return
	}
	defer dial.Close()
	client := product.NewProductServiceClient(dial)
	var response *product.FetchReply
	response, err = client.Fetch(context.TODO(), &product.FetchQuery{Url: args[0]})
	if err != nil {
		 return
	}
	fmt.Printf("Fetched %d rows\n", response.Size)
	return
}

func init() {
	rootCmd.AddCommand(fetchCmd)
}
