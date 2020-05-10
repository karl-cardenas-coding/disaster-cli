package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/cobra/doc"
)

var (
	ApikeyFlag       string
	OutputFlag       string
	DisplayMapFlag   bool
	GenerateDocFlag  bool
	VersionString    string = "No version provided"
	FiltersFlag             = make([]string, 0, 10)
	DownloadTempPath string
)

var rootCmd = &cobra.Command{
	Use:   "disaster",
	Short: "A CLI too for determining natural catastrophe near you, or a location specified",
	Long:  `A Golang based CLI too for determining natural catastrophe near you, or a location specified. Visit https://github.com/karl-cardenas-coding/disaster-cli for more information.`,
	Run: func(cmd *cobra.Command, args []string) {
		generateDocFlag := GenerateDocFlag
		if generateDocFlag {
			err := doc.GenMarkdownTree(cmd, "./documentation/")
			if err != nil {
				log.Fatal(err)
			}
		}

		err := cmd.Help()
		if err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&ApikeyFlag, "api-key", "a", "", "Override default apikey from nasa.gov")
	rootCmd.PersistentFlags().StringVarP(&OutputFlag, "output", "o", "text", "Output format options: table | text | json")
	rootCmd.PersistentFlags().BoolVarP(&GenerateDocFlag, "documentation", "c", false, "Generate documentation")
	eventsCmd.Flags().BoolVarP(&DisplayMapFlag, "display-map", "d", false, "Displays the Google Maps URL")
	eventsCmd.Flags().StringSliceVarP(&FiltersFlag, "filter", "f", []string{}, "filter events by passing in categories (comma seperated")
	updateCmd.Flags().StringVarP(&DownloadTempPath, "temp-location", "l", "", "Specify the temporary directory to use for the update process")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
