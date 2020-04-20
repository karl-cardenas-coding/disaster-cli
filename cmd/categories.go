package cmd

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/jedib0t/go-pretty/table"
	"github.com/karl-cardenas-coding/disaster-cli/library"
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
	Long: `Prints all the unique categories of all the events occuring right now.
The returned category may be used with the --filter flag for the event cmd.`,
	Run: func(cmd *cobra.Command, args []string) {

		apikey := ApikeyFlag
		outputFlag := OutputFlag

		records := library.QueryCategoriesAPI(apikey)

		if outputFlag == "text" {

			for _, v := range records.Categories {

				fmt.Println(v.ID)

			}
			fmt.Printf("\nThere are currently %v unique natural disaster categories occuring in the world.\n\n", len(records.Categories))
		}

		if outputFlag == "table" {
			outputCategoriesTable(records)
		}

		if outputFlag == "json" {
			json, err := json.Marshal(records)
			if err != nil {
				fmt.Println(err)
			}
			os.Stdout.Write(json)
		}

	},
}
