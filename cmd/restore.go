package cmd

import (
	"github.com/jkaninda/s3fs/pkg"
	"github.com/spf13/cobra"
)

var RestoreCmd = &cobra.Command{
	Use:   "restore",
	Short: "Restore file from S3",
	Run: func(cmd *cobra.Command, args []string) {
		pkg.Restore(cmd)
	},
}
