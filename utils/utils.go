package utils

import (
	"fmt"
	"github.com/spf13/cobra"
	"io/fs"
	"os"
)

func Info(v ...any) {
	fmt.Println("⒤ ", fmt.Sprint(v...))
}
func Done(v ...any) {
	fmt.Println("✔ ", fmt.Sprint(v...))
}
func Fatal(v ...any) {
	fmt.Println("✘ ", fmt.Sprint(v...))
	os.Exit(1)
}
func Fatalf(msg string, v ...any) {
	fmt.Printf("✘ "+msg, v...)
	os.Exit(1)
}

func FileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func WriteToFile(filePath, content string) error {
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString(content)
	return err
}
func ChangePermission(filePath string, mod int) {
	if err := os.Chmod(filePath, fs.FileMode(mod)); err != nil {
		Fatalf("Error changing permissions of %s: %v\n", filePath, err)
	}

}
func IsDirEmpty(name string) (bool, error) {
	f, err := os.Open(name)
	if err != nil {
		return false, err
	}
	defer f.Close()

	_, err = f.Readdirnames(1)
	if err == nil {
		return false, nil
	}
	return true, nil
}

func GetEnv(cmd *cobra.Command, flagName, envName string) string {
	value, _ := cmd.Flags().GetString(flagName)
	if value != "" {
		err := os.Setenv(envName, value)
		if err != nil {
			return value
		}
	}
	return os.Getenv(envName)
}
func FlagGetString(cmd *cobra.Command, flagName string) string {
	value, _ := cmd.Flags().GetString(flagName)
	if value != "" {
		return value

	}
	return ""
}
func FlagGetBool(cmd *cobra.Command, flagName string) bool {
	value, _ := cmd.Flags().GetBool(flagName)
	return value
}

func SetEnv(key, value string) {

	err := os.Setenv(key, value)
	if err != nil {
		return
	}
}
