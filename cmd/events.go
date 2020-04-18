package cmd

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/jedib0t/go-pretty/table"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(eventsCmd)
}

func outputTable(records Response) {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"#", "Title", "Description", "Category", "Map Link"})
	for i, v := range records.Events {

		if v.Description == nil {
			t.AppendRow([]interface{}{i + 1, v.Title, "N/A", v.Categories[0].Title})
		}

		if v.Description != nil {
			t.AppendRow([]interface{}{i + 1, v.Title, v.Description, v.Categories[0].Title})
		}

		// fmt.Printf("%v\t Link: https://www.google.com/maps/@?api=1&map_action=map&center=%v,%v&zoom=12&basemap=terrain \n", v.Title, long, lat)
	}

	t.AppendFooter(table.Row{"Total", len(records.Events)})
	// t.SetStyle(table.StyleColoredBright)
	t.Render()
}

type Response struct {
	Title       string   `json:"title"`
	Description string   `json:"description"`
	Link        string   `json:"link"`
	Events      []Events `json:"events"`
}
type Categories struct {
	ID    string `json:"id"`
	Title string `json:"title"`
}
type Sources struct {
	ID  string `json:"id"`
	URL string `json:"url"`
}
type Geometry struct {
	MagnitudeValue float64   `json:"magnitudeValue"`
	MagnitudeUnit  string    `json:"magnitudeUnit"`
	Date           time.Time `json:"date"`
	Type           string    `json:"type"`
	Coordinates    []float64 `json:"coordinates"`
}
type Events struct {
	ID          string        `json:"id"`
	Title       string        `json:"title"`
	Description interface{}   `json:"description"`
	Link        string        `json:"link"`
	Closed      interface{}   `json:"closed"`
	Categories  []Categories  `json:"categories"`
	Sources     []Sources     `json:"sources"`
	Geometry    []interface{} `json:"geometry"`
}

var eventsCmd = &cobra.Command{
	Use:   "events",
	Short: "Returns all events occuring in the world at this point in time.",
	Long:  `Return all defined events in the world`,
	Run: func(cmd *cobra.Command, args []string) {

		// Retrive api key if flag parameter is passed down
		apikey := Apikey
		output := Output

		if apikey == "" {
			apikey = "bhKzxWngdTEBCJLSnmIMLW5KqAjVPqEOKCdqK6Wn"
		}

		url := fmt.Sprintf("https://eonet.sci.gsfc.nasa.gov/api/v3/events?api_key=%s", apikey)

		resp, err := http.Get(url)
		if err != nil {
			fmt.Println("Error:", err)
		}
		defer resp.Body.Close()

		var records Response

		if err := json.NewDecoder(resp.Body).Decode(&records); err != nil {
			log.Println(err)
		}

		// Output flag
		if output == "text" {
			for _, v := range records.Events {
				fmt.Printf("%v\n", v.Title)
			}
			fmt.Printf("\nThere are currently %v natural catastrophe events occuring in the world.\n\n", len(records.Events))
		}

		if output == "table" {
			outputTable(records)
		}
		// Output flag logic end
	},
}
