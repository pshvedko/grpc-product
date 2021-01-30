package cmd

import (
	"fmt"
	"io"

	"github.com/spf13/cobra"
)

var fetchCmd = &cobra.Command{
	Use:   "fetch url",
	Args:  cobra.ExactArgs(1),
	Short: "Load external CSV file",
	Long: `The command loads a CSV file with the 'PRODUCT NAME; PRICE' format
from the specified URL and saves it to MongoDB`,
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("fetch called", args)
		return io.EOF
	},
}

func init() {
	rootCmd.AddCommand(fetchCmd)
}
