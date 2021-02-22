package cmd

import (
	"context"
	"encoding/json"
	"os"

	"github.com/pshvedko/grpc-product/product"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Browses the contents of storage",
	Long:  "FIXME",
	RunE:  runList,
}

func runList(*cobra.Command, []string) (err error) {
	var dial *grpc.ClientConn
	dial, err = grpc.Dial(addressFlag.String(), grpc.WithInsecure())
	if err != nil {
		return
	}
	defer dial.Close()
	client := product.NewProductServiceClient(dial)
	var sort []*product.Sort
	for _, v := range sortByFlag {
		sort = append(sort, product.NewSort(v))
	}
	var reply *product.ListReply
	reply, err = client.List(context.TODO(), &product.ListQuery{
		Page: &product.Page{
			Limit:  limitFlag,
			Offset: offsetFlag,
		},
		Sort: sort,
	})
	j := json.NewEncoder(os.Stdout)
	if err != nil {
		return
	}
	for _, p := range reply.Products {
		err = j.Encode(p)
		if err != nil {
			return
		}
	}
	return
}

var limitFlag uint32
var offsetFlag uint32
var sortByFlag []string

func init() {
	listCmd.PersistentFlags().Uint32VarP(&limitFlag, "limit", "l", 0, "limit")
	listCmd.PersistentFlags().Uint32VarP(&offsetFlag, "offset", "o", 0, "offset")
	listCmd.PersistentFlags().StringArrayVarP(&sortByFlag, "sort", "s", nil, "field [+/-]name(s) to sort")
	rootCmd.AddCommand(listCmd)
}
