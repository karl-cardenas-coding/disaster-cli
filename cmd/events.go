package cmd

import (
  "fmt"
  "net/http"
  "log"
  "time"
  "encoding/json"

  "github.com/spf13/cobra"
)

func init() {
 rootCmd.AddCommand(eventsCmd)
}

type Events struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Link        string `json:"link"`
	Events      []struct {
		ID          string `json:"id"`
		Title       string `json:"title"`
		Description string `json:"description"`
		Link        string `json:"link"`
		Categories  []struct {
			ID    int    `json:"id"`
			Title string `json:"title"`
		} `json:"categories"`
		Sources []struct {
			ID  string `json:"id"`
			URL string `json:"url"`
		} `json:"sources"`
		Geometries []struct {
			Date        time.Time `json:"date"`
			Type        string    `json:"type"`
			Coordinates []interface{} `json:"coordinates"`
		} `json:"geometries"`
	} `json:"events"`
}

var eventsCmd = &cobra.Command{
  Use:   "events",
  Short: "Returns all events occuring in the world at this point in time.",
  Long:  `Return all defined events in the world`,
  Run: func(cmd *cobra.Command, args []string) {
    //TODO add code to query API
  url := "https://eonet.sci.gsfc.nasa.gov/api/v2.1/events?api_key=bhKzxWngdTEBCJLSnmIMLW5KqAjVPqEOKCdqK6Wn"

  resp, err := http.Get(url)
  if err != nil {
  	fmt.Println("Error:", err)
  }
  defer resp.Body.Close()

  var records Events

  if err := json.NewDecoder(resp.Body).Decode(&records); err != nil {
		log.Println(err)
	}

  fmt.Printf("There are currently %v natural catastrophe events occuring in the world.\n\n", len(records.Events))
  for _,v := range records.Events {
    // var lat  float64
    // var long float64

    // for _,v2 := range v.Geometries {
    //    lat = v2.Coordinates[0].(float64)
    //    long = v2.Coordinates[1].(float64)
    // }
    fmt.Printf("%v\t %v\n",v.Title, v.Description)
    // fmt.Printf("%v\t Link: https://www.google.com/maps/@?api=1&map_action=map&center=%v,%v&zoom=12&basemap=terrain \n", v.Title, long, lat)
  }

  // fmt.Println(records)
  },
}
