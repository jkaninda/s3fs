// Package pkg /*
/*
Copyright © 2024 Jonas Kaninda
*/
package pkg

import (
	"fmt"
	"github.com/jkaninda/s3fs/utils"
	"github.com/spf13/cobra"
	"os"
	"os/exec"
)

// MountS3Storage Mount s3 storage using s3fs
func MountS3Storage(cmd *cobra.Command) {

	keep, _ = cmd.Flags().GetBool("keep")

	accessKey = os.Getenv("ACCESS_KEY")
	secretKey = os.Getenv("SECRET_KEY")
	bucketName = os.Getenv("BUCKET_NAME")
	if bucketName == "" {
		bucketName = os.Getenv("BUCKETNAME")
	}
	s3Endpoint = os.Getenv("S3_ENDPOINT")

	if accessKey == "" || secretKey == "" || bucketName == "" {
		utils.Fatal("Please make sure all environment variables are set for S3")
	} else {
		storagePath := s3MountPath

		//Write file
		err := utils.WriteToFile(s3fsPasswdFile, fmt.Sprintf("%s:%s", accessKey, secretKey))
		if err != nil {
			utils.Fatal("Error creating file")
		}
		//Change file permission
		utils.ChangePermission(s3fsPasswdFile, 0600)

		//Mount object storage
		utils.Info("Mounting Object storage in ", s3MountPath)
		if isEmpty, _ := utils.IsDirEmpty(s3MountPath); isEmpty {
			cmd := exec.Command("s3fs", bucketName, s3MountPath,
				"-o", "passwd_file="+s3fsPasswdFile,
				"-o", "use_cache=/tmp/s3cache",
				"-o", "allow_other",
				"-o", "url="+s3Endpoint,
				"-o", "use_path_request_style",
			)

			if err := cmd.Run(); err != nil {
				utils.Fatal("Error mounting Object storage:", err)
			}
			utils.Info("Listing directory")
			//listDir()

			if keep {
				//Start Supervisor
				utils.Info("Running in background")
				supervisordCmd := exec.Command("supervisord", "-c", "/etc/supervisor/supervisord.conf")
				if err := supervisordCmd.Run(); err != nil {
					utils.Fatalf("Error starting supervisord: %v\n", err)
				}
			}

		} else {
			utils.Info("Object storage already mounted in " + s3MountPath)
			if err := os.MkdirAll(storagePath, os.ModePerm); err != nil {
				utils.Fatal("Error creating directory "+storagePath, err)
			}

		}

	}
}
func Backup(cmd *cobra.Command) {

}
func Restore(cmd *cobra.Command) {

}
func listDir() {
	entries, err := os.ReadDir(s3MountPath)
	if err != nil {
		utils.Fatal("Error listing files")
	}

	for _, e := range entries {
		fmt.Println(e.Name())
	}
}
