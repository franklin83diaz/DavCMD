package cmd

import (
	"devcmd/internal/webdav"

	"github.com/spf13/cobra"
)

var uploadCmd = &cobra.Command{
	Use:   "upload",
	Short: "upload file",
	Long:  `upload file to webdav server`,
	Run: func(cmd *cobra.Command, args []string) {
		if directory {
			if zip {
				webdav.UploadZip(url, args[0], args[1], username, password)
			}
		} else {
			webdav.Upload(url, args[0], args[1], username, password, zip)
		}
	},
}
