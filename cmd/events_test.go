package cmd

import (
	"fmt"
	// "io/ioutil"
	"os"
	"testing"

	"github.com/spf13/cobra"
)

func versionTestCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "version",
		Short: "Print the version number of disaster-cli",
		Long:  `Prints the versio number of disaster-cli`,
		Run: func(cmd *cobra.Command, args []string) {
			os.Stdout.Write([]byte("disaster-cli v0.1"))
		},
	}
}

func Test_ExecuteCommand(t *testing.T) {
	cmd := versionTestCmd()
	cmd.Execute()
	r, w, _ := os.Pipe()
	os.Stdout = w
	w.Close()
}

// func isWindowsCI() bool {
// 	return runtime.GOOS == "windows" && os.Getenv("CI") != ""
// }
