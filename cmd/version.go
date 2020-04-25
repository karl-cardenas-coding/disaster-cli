package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the current version number of disaster-cli",
	Long:  `Prints the current version number of disaster-cli`,
	Run: func(cmd *cobra.Command, args []string) {
		os.Stdout.Write([]byte("disaster-cli v1.0.0"))
	},
}
