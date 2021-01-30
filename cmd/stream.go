package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// streamCmd represents the stream command
var streamCmd = &cobra.Command{
	Use:   "stream",
	Short: "A brief description of your command",
	Long:  "FIXME",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("stream called")
	},
}

func init() {
	listCmd.AddCommand(streamCmd)
}
