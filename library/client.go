package library

import (
	"encoding/json"
	"net/http"
	"time"

	log "github.com/sirupsen/logrus"
)

const (
	ISSUE_MSG = " Please open up a Github issue to report this error! https://github.com/karl-cardenas-coding/disaster-cli"
)

var disasterClient = &http.Client{
	Timeout: 10 * time.Second,
	Transport: &http.Transport{
		MaxIdleConns:          10,
		ResponseHeaderTimeout: 10 * time.Second,
		IdleConnTimeout:       5 * time.Second,
		DisableCompression:    true,
		ForceAttemptHTTP2:     true,
	},
}

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
