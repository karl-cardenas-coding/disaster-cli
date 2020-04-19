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
	// var DisplayMap bool
	rootCmd.AddCommand(eventsCmd)
	// eventsCmd.Flags().BoolVarP(&DisplayMap, "display-map", "dm", false, "Displays the Google Maps URL")
}

// Function for displaying the data in a table
func outputTable(records library.Response) {
	// var lat, long float64

	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"#", "Title", "Category"})
	for i, v := range records.Events {

		// lat, long = getLocation(v.Geometry[0].(map[string]interface{}))
		// mapLink := fmt.Sprintf("https://www.google.com/maps/@?api=1&map_action=map&center=%v,%v&zoom=6&basemap=terrain", long, lat)

		t.AppendRow([]interface{}{i + 1, v.Title, v.Categories[0].Title})
	}

	t.AppendFooter(table.Row{"Total", len(records.Events)})
	// t.SetStyle(table.StyleColoredBright)
	t.Render()
}

// Function for retriving the geograhical location
func getLocation(input map[string]interface{}) (float64, float64) {
	var lat, long float64
	temp := input["coordinates"].([]interface{})

	_, ok := temp[0].(float64)
	if ok {
		lat = temp[0].(float64)
		long = temp[1].(float64)
	}
	return lat, long
}

var eventsCmd = &cobra.Command{
	Use:   "events",
	Short: "Returns all events occuring in the world at this point in time.",
	Long:  `Return all defined events in the world`,
	Args:  cobra.MaximumNArgs(0),
	Run: func(cmd *cobra.Command, args []string) {

		// Retrive api key if flag parameter is passed down
		apikey := ApikeyFlag
		outputFlag := OutputFlag
		displayMapFlag := DisplayMapFlag

		records := library.QueryAPI(apikey)

		// Output flag
		if outputFlag == "text" {

			for _, v := range records.Events {

				// URL to form Google Map link:  https://www.google.com/maps/@?api=1&map_action=map&center=-37.29356,-71.95059&zoom=6&basemap=terrain
				// The coordinates received from the events payload need to be revered in the URL:
				// https://eonet.sci.gsfc.nasa.gov/api/v3/events
				var lat, long float64
				lat, long = getLocation(v.Geometry[0].(map[string]interface{}))
				mapLink := fmt.Sprintf("https://www.google.com/maps/@?api=1&map_action=map&center=%v,%v&zoom=6&basemap=terrain", long, lat)

				if displayMapFlag {
					fmt.Printf("%v\nMap Link: %v\n--------------------------------\n", v.Title, mapLink)
				} else {
					fmt.Printf("%v\n", v.Title)

				}

			}
			fmt.Printf("\nThere are currently %v natural catastrophes events occuring in the world.\n\n", len(records.Events))
		}

		if outputFlag == "table" {
			outputTable(records)
		}

		if outputFlag == "json" {
			json, err := json.Marshal(records)
			if err != nil {
				fmt.Println(err)
			}
			os.Stdout.Write(json)
		}
		// Output flag logic end
	},
}
