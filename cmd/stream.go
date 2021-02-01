package cmd

import (
	"github.com/spf13/cobra"
)

var streamCmd = &cobra.Command{
	Use:   "stream",
	Short: "A brief description of your command",
	Long:  "FIXME",
	RunE:  runStream,
}

func runStream(*cobra.Command, []string) error {
	return nil
}

func init() {
	listCmd.AddCommand(streamCmd)
}
