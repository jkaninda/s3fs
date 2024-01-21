package cmd

import (
	"github.com/jkaninda/s3fs/pkg"
	"github.com/spf13/cobra"
)

var S3MountCmd = &cobra.Command{
	Use:   "mount",
	Short: "Mount AWS S3 storage",
	Run: func(cmd *cobra.Command, args []string) {
		pkg.MountS3Storage(cmd)
	},
}

func init() {
	S3MountCmd.PersistentFlags().BoolP("keep", "k", false, "Keep program running in backup")

}
