// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MIT

package library

import (
	"encoding/json"
	"net/http"

	log "github.com/sirupsen/logrus"
)

const (
	ISSUE_MSG = " Please open up a Github issue to report this error! https://github.com/karl-cardenas-coding/disaster-cli"
)

func getJson(url string, target interface{}) error {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Println("unable to create HTTP request")
	}

	req.Header = http.Header{
		"Content-Type":    []string{`application/json; charset=utf-8`},
		"Accept-Encoding": []string{"gzip, deflate, br"},
		"Accept":          []string{"application/json"},
	}

	disasterClient := &http.Client{
		Transport: &http.Transport{
			ForceAttemptHTTP2: true,
		},
	}

	r, err := disasterClient.Do(req)
	if err != nil {
		log.WithFields(log.Fields{
			"package":         "cmd",
			"file":            "client.go",
			"parent_function": "getJSON",
			"function":        "disasterClient.Get",
			"error":           err,
			"data":            url,
		}).Error("Error conencting to NASA's API.", ISSUE_MSG)
		return err
	}
	defer r.Body.Close()

	return json.NewDecoder(r.Body).Decode(target)
}
