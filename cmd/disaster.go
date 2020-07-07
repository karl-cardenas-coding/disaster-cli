package cmd

import (
	"os"

	log "github.com/sirupsen/logrus"

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

const (
	ISSUE_MSG = " Please open up a Github issue to report this error! https://github.com/karl-cardenas-coding/disaster-cli"
)

var rootCmd = &cobra.Command{
	Use:   "disaster",
	Short: "A CLI tool for determining natural catastrophe near you, or a location specified",
	Long:  `A Golang based CLI tool for determining natural catastrophe near you, or a location specified. Visit https://github.com/karl-cardenas-coding/disaster-cli for more information.`,
	Run: func(cmd *cobra.Command, args []string) {
		generateDocFlag := GenerateDocFlag
		if generateDocFlag {
			err := doc.GenMarkdownTree(cmd, "./documentation/")
			if err != nil {
				log.WithFields(log.Fields{
					"package":  "cmd",
					"file":     "disaster.go",
					"parent_function": "generateDocFlag",
					"function": "doc.GenMarkdownTree"
					"error":    err,
					"data":     fmt.Sprint(cmd, "./documentation/"),
				}).Fatal("Error generating markdown content", ISSUE_MSG)
			}
		}

		err := cmd.Help()
		if err != nil {
			log.WithFields(log.Fields{
				"package":  "cmd",
				"file":     "disaster.go",
				"parent_function": "generateDocFlag",
				"function": "cmd.Help",
				"error":    err,
				"data":     nil,
			}).Fatal("Error outputting help!", ISSUE_MSG)
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

	// Establish logging default
	log.SetFormatter(&log.TextFormatter{
		DisableColors: false,
		FullTimestamp: true,
	})
	log.SetOutput(os.Stdout)

	log.SetLevel(log.WarnLevel)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.WithFields(log.Fields{
			"package":  "cmd",
			"file":     "disaster.go",
			"function": "Execute",
			"error":    err,
			"data":     nil,
		}).Fatal("Error executing the CLI!", ISSUE_MSG)
		os.Exit(1)
	}
}
