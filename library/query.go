package library

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func QueryEventAPI(apikey string) EventResponse {
	url := fmt.Sprintf("https://eonet.sci.gsfc.nasa.gov/api/v3/events?api_key=%s", apikey)

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error:", err)
	}
	defer resp.Body.Close()

	var records EventResponse

	if err := json.NewDecoder(resp.Body).Decode(&records); err != nil {
		log.Println(err)
	}

	return records

}

func QueryCategoriesAPI(apikey string) CategoriesResponse {
	url := fmt.Sprintf("https://eonet.sci.gsfc.nasa.gov/api/v3/categories?api_key=%s", apikey)

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error:", err)
	}
	defer resp.Body.Close()

	var records CategoriesResponse

	if err := json.NewDecoder(resp.Body).Decode(&records); err != nil {
		log.Println(err)
	}

	return records

}
