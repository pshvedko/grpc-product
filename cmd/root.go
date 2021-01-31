package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

var portFlag uint16

var rootCmd = &cobra.Command{
	Use: filepath.Base(os.Args[0]),
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().Uint16VarP(&portFlag, "port", "p", 9000, "Port to listen")
}
