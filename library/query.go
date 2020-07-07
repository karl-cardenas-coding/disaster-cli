package library

import (
	"encoding/json"
	"fmt"
	"net/http"

	log "github.com/sirupsen/logrus"
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

func QueryCategoriesAPI(apikey string, category string) CategoriesResponse {
	var url string

	if category == "" {
		url = fmt.Sprintf("https://eonet.sci.gsfc.nasa.gov/api/v3/categories/%s?api_key=%s", category, apikey)
	} else {
		url = fmt.Sprintf("https://eonet.sci.gsfc.nasa.gov/api/v3/categories?api_key=%s", apikey)

	}

	resp, err := http.Get(url)
	if err != nil {
		log.WithFields(log.Fields{
			"package":         "cmd",
			"file":            "query.go",
			"parent_function": "QueryCategoriesAPI",
			"function":        "http.Get",
			"error":           err,
			"data":            url,
		}).Error("Error conencting to NASA's API.", ISSUE_MSG)
	}
	defer resp.Body.Close()

	var records CategoriesResponse

	if err := json.NewDecoder(resp.Body).Decode(&records); err != nil {
		log.WithFields(log.Fields{
			"package":         "cmd",
			"file":            "query.go",
			"parent_function": "QueryCategoriesAPI",
			"function":        "json.NewDecoder",
			"error":           err,
			"data":            records,
		}).Error("Error conencting to NASA's API.", ISSUE_MSG)
	}

	return records

}
