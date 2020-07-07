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

var disasterClient = &http.Client{Timeout: 10 * time.Second}

func getJson(url string, target interface{}) error {
	r, err := disasterClient.Get(url)
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
