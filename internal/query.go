// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MIT

package library

import (
	"encoding/json"
	"fmt"
	"net/http"

	log "github.com/sirupsen/logrus"
)

func QueryEventAPI(apikey string) EventResponse {

	var (
		url     string
		records EventResponse
	)
	if apikey != "" {
		url = fmt.Sprintf("https://eonet.gsfc.nasa.gov/api/v3/events?api_key=%s", apikey)
	} else {
		url = "https://eonet.gsfc.nasa.gov/api/v3/events"
	}

	client := &http.Client{
		Transport: &http.Transport{
			ForceAttemptHTTP2: true,
		},
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Println("unable to create HTTP request")
	}

	req.Header = http.Header{
		"Content-Type":    []string{`application/json; charset=utf-8`},
		"Accept-Encoding": []string{"gzip, deflate, br"},
		"Accept":          []string{"application/json"},
	}

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error:", err)
	}
	defer resp.Body.Close()

	if err := json.NewDecoder(resp.Body).Decode(&records); err != nil {
		log.Println(err)
	}

	return records

}

func QueryCategoriesAPI(apikey string, category string) CategoriesResponse {
	var (
		url     string
		records CategoriesResponse
	)

	if category == "" {
		url = fmt.Sprintf("https://eonet.gsfc.nasa.gov/api/v3/categories/%s?api_key=%s", category, apikey)
	} else {
		url = fmt.Sprintf("https://eonet.gsfc.nasa.gov/api/v3/categories?api_key=%s", apikey)

	}

	client := &http.Client{
		Transport: &http.Transport{
			ForceAttemptHTTP2: true,
		},
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Println("unable to create HTTP request")
	}

	req.Header = http.Header{
		"Content-Type":    []string{`application/json; charset=utf-8`},
		"Accept-Encoding": []string{"gzip, deflate, br"},
		"Accept":          []string{"application/json"},
	}

	resp, err := client.Do(req)
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
