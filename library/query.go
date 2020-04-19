package library

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func QueryAPI(apikey string) Response {
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

	return records

}
