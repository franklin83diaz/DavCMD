package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Show DavCMD version",
	Long:  `Show DavCMD version and exit`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("DavCMD v0.0.1 Beta")

	},
}
