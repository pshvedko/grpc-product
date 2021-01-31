package cmd

import (
	"fmt"
	"net"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

import _ "github.com/pshvedko/grpc-product/migrate"

var portFlag uint16
var addrFlag net.IP

var rootCmd = &cobra.Command{
	Use:           filepath.Base(os.Args[0]),
	SilenceErrors: true,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().Uint16VarP(&portFlag, "port", "p", 9000, "port to listen")
	rootCmd.PersistentFlags().IPVarP(&addrFlag, "addr", "a", net.IP{0, 0, 0, 0}, "address to bind")
}
