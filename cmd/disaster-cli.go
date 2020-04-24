package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/cobra/doc"
)

var ApikeyFlag string
var OutputFlag string
var DisplayMapFlag bool

// var FilterFlag []string

var rootCmd = &cobra.Command{
	Use:   "disaster",
	Short: "A CLI too for determining natural catastrophe near you, or a location specified",
	Long:  `A Golang based CLI too for determining natural catastrophe near you, or a location specified. Visit https://github.com/karl-cardenas-coding/disaster-cli for more information.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(`
A Golang based CLI too for determining natural catastrophe near you, or a location specified. Visit https://github.com/karl-cardenas-coding/disaster-cli for more information.

Usage:
  disaster [flags]
  disaster [command]

Available Commands:
  categories  Prints all the unique categories of all the events
  events      Returns all events occurring in the world at this point in time.
  help        Help about any command
  version     Print the version number of disaster-cli

Flags:
  -a, --api-key string   Override default apikey from nasa.gov
  -h, --help             help for disaster-cli
  -o, --output string    Output format options: table | text | json (default "text")

Use "disaster [command] --help" for more information about a command.`)
		err := doc.GenMarkdownTree(cmd, "./documentation/")
		if err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&ApikeyFlag, "api-key", "a", "", "Override default apikey from nasa.gov")
	rootCmd.PersistentFlags().StringVarP(&OutputFlag, "output", "o", "text", "Output format options: table | text | json")
	eventsCmd.Flags().BoolVarP(&DisplayMapFlag, "display-map", "d", false, "Displays the Google Maps URL")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
