package cmd

import (
	"github.com/jkaninda/s3fs/pkg"
	"github.com/spf13/cobra"
)

var BackupCmd = &cobra.Command{
	Use:   "backup",
	Short: "backup file to S3",
	Run: func(cmd *cobra.Command, args []string) {
		pkg.Backup(cmd)
	},
}

func init() {
	S3MountCmd.PersistentFlags().BoolP("enable-compression", "", false, "Enable backup folder compression")

}
