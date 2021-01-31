package cmd

import (
	"github.com/spf13/cobra"
)

//c := product.NewProductServiceClient(dial)
//var reply *product.ListReply
//reply, err = c.List(context.TODO(), &product.ListQuery{Page: &product.Page{
//Limit:  10,
//Offset: 20,
//}, Sort: []*product.Sort{{
//Order: false,
//By:    "price",
//}, {
//Order: true,
//By:    "name",
//}}})
//if err != nil {
//log.Fatal(err)
//}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "A brief description of your command",
	Long:  "FIXME",
	RunE:  runList,
}

func runList(*cobra.Command, []string) error {
	return nil
}

var limitFlag uint32
var offsetFlag uint32

func init() {
	listCmd.PersistentFlags().Uint32VarP(&limitFlag, "limit", "l", 0, "limit rows on page")
	listCmd.PersistentFlags().Uint32VarP(&offsetFlag, "offset", "0", 0, "offset of page")
	rootCmd.AddCommand(listCmd)
}
