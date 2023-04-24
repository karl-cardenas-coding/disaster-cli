// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MIT

package cmd

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/karl-cardenas-coding/disaster-cli/library"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(eventsCmd)
}

// Function for displaying the data in a table
func outputTable(records library.EventResponse, filters []string) {
	// var lat, long float64
	counter := 0

	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"#", "Title", "Category"})

	if len(filters) > 0 {

		for _, v := range records.Events {

			for _, filter := range filters {

				if filter == v.Categories[0].ID {
					// lat, long = getLocation(v.Geometry[0].(map[string]interface{}))
					// mapLink := fmt.Sprintf("https://www.google.com/maps/@?api=1&map_action=map&center=%v,%v&zoom=6&basemap=terrain", long, lat)
					counter++
					t.AppendRow([]interface{}{counter, v.Title, v.Categories[0].Title})
				}
			}

		}

	}

	if len(filters) == 0 {
		for _, v := range records.Events {
			counter++

			// lat, long = getLocation(v.Geometry[0].(map[string]interface{}))
			// mapLink := fmt.Sprintf("https://www.google.com/maps/@?api=1&map_action=map&center=%v,%v&zoom=6&basemap=terrain", long, lat)

			t.AppendRow([]interface{}{counter, v.Title, v.Categories[0].Title})
		}
	}

	t.AppendFooter(table.Row{"Total", counter})
	// t.SetStyle(table.StyleColoredBright)
	t.Render()
}

func outputJSON(records library.EventResponse, filters []string) {

	if len(filters) > 0 {
		var list []library.Events

		for _, v := range records.Events {

			for _, filter := range filters {

				if filter == v.Categories[0].ID {
					list = append(list, v)
				}
			}

		}

		json, err := json.MarshalIndent(&list, " ", " ")
		if err != nil {
			log.WithFields(log.Fields{
				"package":         "cmd",
				"file":            "events.go",
				"parent_function": "outputJSON",
				"function":        "json.Marshal",
				"error":           err,
				"data":            list,
			}).Error("Error marshalling JSON", ISSUE_MSG)
		}
		os.Stdout.Write(json)

	}

	if len(filters) == 0 {

		json, err := json.MarshalIndent(&records, " ", " ")
		if err != nil {
			log.WithFields(log.Fields{
				"package":         "cmd",
				"file":            "events.go",
				"parent_function": "outputJSON",
				"function":        "json.Marshal",
				"error":           err,
				"data":            records,
			}).Error("Error marshalling JSON", ISSUE_MSG)
		}
		os.Stdout.Write(json)

	}

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

func outputText(records library.EventResponse, displayMap bool, filters []string) {
	counter := 0

	for _, v := range records.Events {

		if len(filters) != 0 {

			for _, filter := range filters {

				if filter == v.Categories[0].ID {
					// URL to form Google Map link:  https://www.google.com/maps/@?api=1&map_action=map&center=-37.29356,-71.95059&zoom=6&basemap=terrain
					// The coordinates received from the events payload need to be revered in the URL:
					// https://eonet.sci.gsfc.nasa.gov/api/v3/events
					var lat, long float64
					lat, long = getLocation(v.Geometry[0].(map[string]interface{}))
					mapLink := fmt.Sprintf("https://www.google.com/maps/@?api=1&map_action=map&center=%v,%v&zoom=6&basemap=terrain", long, lat)
					counter++

					if displayMap {
						fmt.Printf("%v\nMap Link: %v\n--------------------------------\n", v.Title, mapLink)
					} else {
						fmt.Printf("%v\n", v.Title)

					}
				}
			}

		}

		if len(filters) == 0 {
			counter++
			// URL to form Google Map link:  https://www.google.com/maps/@?api=1&map_action=map&center=-37.29356,-71.95059&zoom=6&basemap=terrain
			// The coordinates received from the events payload need to be revered in the URL:
			// https://eonet.sci.gsfc.nasa.gov/api/v3/events
			var lat, long float64
			lat, long = getLocation(v.Geometry[0].(map[string]interface{}))
			mapLink := fmt.Sprintf("https://www.google.com/maps/@?api=1&map_action=map&center=%v,%v&zoom=6&basemap=terrain", long, lat)

			if displayMap {
				fmt.Printf("%v\nMap Link: %v\n--------------------------------\n", v.Title, mapLink)
			} else {
				fmt.Printf("%v\n", v.Title)
			}
		}

	}
	fmt.Printf("\nThere are currently %v natural catastrophes events occuring in the world.\n\n", counter)
}

var eventsCmd = &cobra.Command{
	Use:   "events",
	Short: "Returns all events occurring in the world at this point in time.",
	Long:  `Return all defined events in the world`,
	Run: func(cmd *cobra.Command, args []string) {

		// Retrive api key if flag parameter is passed down
		// Initialize parameters
		apikey := ApikeyFlag
		outputFlag := OutputFlag
		displayMapFlag := DisplayMapFlag
		filtersFlags := FiltersFlag

		records := library.QueryEventAPI(apikey)

		if outputFlag == "text" {
			outputText(records, displayMapFlag, filtersFlags)
		}

		if outputFlag == "table" {
			outputTable(records, filtersFlags)
		}

		if outputFlag == "json" {
			outputJSON(records, filtersFlags)
		}
	},
}
