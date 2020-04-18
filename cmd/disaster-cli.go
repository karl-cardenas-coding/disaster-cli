package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var Apikey string
var Output string

var rootCmd = &cobra.Command{
	Use:   "disaster-cli",
	Short: "A CLI too for determining natural catastrophe near you, or a location specified",
	Long:  `A Golang based CLI too for determining natural catastrophe near you, or a location specified. Visit https://github.com/karl-cardenas-coding/disaster-cli for more information.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(`No arguments passed to disaster-cli.
Please issue "disaster-cli help" for further guidance.`)
	},
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&Apikey, "api-key", "a", "", "Override default apikey from nasa.gov")
	rootCmd.PersistentFlags().StringVarP(&Output, "output", "o", "table", "Output formats options: table | text")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
