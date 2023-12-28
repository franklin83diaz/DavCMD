package cmd

import (
	"devcmd/internal/webdav"

	"github.com/spf13/cobra"
)

var downloadCmd = &cobra.Command{
	Use:   "download",
	Short: "download file",
	Long:  `download file to webdav server`,
	Run: func(cmd *cobra.Command, args []string) {
		webdav.DownloadFile(url, username, password, args[0])
	},
}
