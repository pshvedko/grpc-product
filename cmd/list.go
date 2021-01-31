package cmd

import (
	"fmt"

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
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("list called")
	},
}

func init() {
	rootCmd.AddCommand(listCmd)

	listCmd.PersistentFlags().String("foo", "", "A help for foo")
}
