package cmd

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"runtime"
	"strings"

	"github.com/karl-cardenas-coding/disaster-cli/library"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(updateCmd)
}

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Updates the local version disaster-cli",
	Long:  `Queries Github to check for a new version, if there is a new version it will install it`,
	Run: func(cmd *cobra.Command, args []string) {
		updateDisasterCli()
	},
}

// A function that will query the Github repo and return the correct download url from the Release Assets
// The download url will correspond to the local operation system
func getReleaseURL() (string, error) {
	url := "https://api.github.com/repos/karl-cardenas-coding/disaster-cli/releases/latest"
	var output string

	client := &http.Client{}
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("Accept", "application/vnd.github.v3+json")

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error:", err)
	}
	defer resp.Body.Close()

	var release library.Release

	// Unmarshal the JSON to the Github Release strcut
	if err := json.NewDecoder(resp.Body).Decode(&release); err != nil {
		log.Println(err)
	}

	// Loop through the assets list and gather the proper zip file url
	for _, v := range release.Assets {
		// Check the string through the determineOS() to identify the proper zip file for the local OS
		if strings.Contains(v.BrowserDownloadURL, determineOS()) {
			output = v.BrowserDownloadURL
		}
	}

	return output, nil

}

func determineOS() string {

	var output string

	switch runtime.GOOS {
	case "windows":
		if runtime.GOARCH == "386" {
			output = fmt.Sprintf("%s-%s", runtime.GOOS, runtime.GOARCH)
		}

		if runtime.GOARCH == "amd64" {
			output = fmt.Sprintf("%s-%s", runtime.GOOS, runtime.GOARCH)
		}

	case "linux":
		if runtime.GOARCH == "386" {
			output = fmt.Sprintf("%s-%s", runtime.GOOS, runtime.GOARCH)
		}

		if runtime.GOARCH == "amd64" {
			output = fmt.Sprintf("%s-%s", runtime.GOOS, runtime.GOARCH)
		}

	case "darwin":
		if runtime.GOARCH == "386" {
			output = fmt.Sprintf("%s-%s", runtime.GOOS, runtime.GOARCH)
		}

		if runtime.GOARCH == "amd64" {
			output = fmt.Sprintf("%s-%s", runtime.GOOS, runtime.GOARCH)
		}

	default:
		panic("Unable to determine OS")
	}

	return output
}

downloadReleaseZip(url) {
// https://golangcode.com/download-a-file-with-progress/
}

func updateDisasterCli() {

	homePath, err := os.UserConfigDir()
	if err != nil {
		fmt.Println("Error")
	}
	fmt.Println(homePath)

	downloadUrl, err := getReleaseURL()
	if err != nil {
		fmt.Println(err)
	}

	downloadReleaseZip(downloadUrl)
}
