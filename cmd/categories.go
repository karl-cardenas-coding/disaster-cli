package cmd

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/jedib0t/go-pretty/table"
	"github.com/karl-cardenas-coding/disaster-cli/library"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(categoriesCmd)
}

func outputCategoriesTable(records library.CategoriesResponse) {

	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"#", "Category"})
	for i, v := range records.Categories {

		t.AppendRow([]interface{}{i + 1, v.ID})
	}

	t.AppendFooter(table.Row{"Total", len(records.Categories)})
	// t.SetStyle(table.StyleColoredBright)
	t.Render()
}

var categoriesCmd = &cobra.Command{
	Use:   "categories",
	Short: "Prints all the unique categories of all the events",
	Long: `Prints all the unique categories of all the events occurring right now.
The returned category may be used with the --filter flag for the event cmd.`,
	Run: func(cmd *cobra.Command, args []string) {

		apikey := ApikeyFlag
		outputFlag := OutputFlag

		records := library.QueryCategoriesAPI(apikey, "")

		if outputFlag == "text" {

			for _, v := range records.Categories {

				fmt.Println(v.ID)

			}
			fmt.Printf("\nThere are %v unique natural disaster categories available in the EONET system.\n\n", len(records.Categories))
			fmt.Println("Use these provided categories with the events command and the filter flag -f \nExample: disaster events -f floods,wildfires")
		}

		if outputFlag == "table" {
			outputCategoriesTable(records)
		}

		if outputFlag == "json" {
			json, err := json.MarshalIndent(records, " ", " ")
			if err != nil {
				log.WithFields(log.Fields{
					"package":         "cmd",
					"file":            "categories.go",
					"parent_function": "Run",
					"function":        "json.Marshal",
					"error":           err,
					"data":            fmt.Sprint(cmd, "./documentation/"),
				}).Error("Error marshaling JSON", ISSUE_MSG)
			}
			os.Stdout.Write(json)
		}

	},
}
