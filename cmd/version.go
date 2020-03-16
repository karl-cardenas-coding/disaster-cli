package cmd

import (
  "fmt"

  "github.com/spf13/cobra"
)

func init() {
  rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
  Use:   "version",
  Short: "Print the version number of disaster-cli",
  Long:  `Prints the versio number of disaster-cli`,
  Run: func(cmd *cobra.Command, args []string) {
    fmt.Println("disaster-cli v0.1")
  },
}
