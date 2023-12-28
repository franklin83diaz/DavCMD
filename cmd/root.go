package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "davcmd",
	Short: "davcmd is a WebDAV client for the command line",
	Long: `
	DavCMD is a lightweight and powerful command line client
	designed to interact with WebDAV servers. It allows users to upload,
	download on a WebDAV server in an efficient and automated manner.`,
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
