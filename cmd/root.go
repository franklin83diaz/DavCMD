package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var username string
var password string = "null"
var url string
var directory bool
var zip bool

var rootCmd = &cobra.Command{
	Use:   "davcmd",
	Short: "davcmd is a WebDAV client for the command line",
	Long: `
	DavCMD is a lightweight and powerful command line client
	designed to interact with WebDAV servers. It allows users to upload,
	download on a WebDAV server in an efficient and automated manner.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("DavCMD v0.0.1 Beta")
		fmt.Println("run 'davcmd --help' for more information")
	},
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&username, "username", "u", "", "username for authentication")
	rootCmd.PersistentFlags().StringVarP(&password, "password", "p", "", "password for authentication")
	rootCmd.PersistentFlags().StringVarP(&url, "url", "l", "", "url of the WebDAV server")
	rootCmd.PersistentFlags().BoolVarP(&directory, "directory", "d", false, "upload a directory")
	rootCmd.PersistentFlags().BoolVarP(&zip, "zip", "z", false, "zip the directory before uploading")

	rootCmd.MarkPersistentFlagRequired("username")
	rootCmd.MarkPersistentFlagRequired("url")

	//commands
	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(uploadCmd)
	rootCmd.AddCommand(downloadCmd)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
