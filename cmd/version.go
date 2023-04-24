// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MIT

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
	Short: "Print the current version number of disaster-cli",
	Long:  `Prints the current version number of disaster-cli`,
	Run: func(cmd *cobra.Command, args []string) {
		version := fmt.Sprintf("disaster %s%s", "v", VersionString)
		fmt.Println(version)
	},
}
