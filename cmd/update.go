package cmd

import (
	"archive/zip"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"runtime"
	"strings"

	"github.com/dustin/go-humanize"
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

func checkForNewRelease() (bool, error) {
	url := "https://api.github.com/repos/karl-cardenas-coding/disaster-cli/releases/latest"
	version := VersionString
	fmt.Println("Current verison is: ", version)
	var output bool = false

	fmt.Println("Checking for new release")
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

	// Check to see if the current version is equivalent to the latest release
	if version != release.TagName {
		fmt.Println("New version available - ", release.TagName)
		output = true
	}

	return output, nil
}

// A function that will query the Github repo and return the correct download url from the Release Assets
// The download url will correspond to the local operation system
func getReleaseURL() (string, string, error) {
	url := "https://api.github.com/repos/karl-cardenas-coding/disaster-cli/releases/latest"
	var outputUrl string
	var outFileName string

	fmt.Println("Querying Github for release details....")
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
			outFileName = v.Name
			outputUrl = v.BrowserDownloadURL
		}
	}

	return outputUrl, outFileName, err

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

//
// https://golangcode.com/download-a-file-with-progress/
// WriteCounter counts the number of bytes written to it. It implements to the io.Writer interface
// and we can pass this into io.TeeReader() which will report progress on each write cycle.
type WriteCounter struct {
	Total uint64
}

func (wc *WriteCounter) Write(p []byte) (int, error) {
	n := len(p)
	wc.Total += uint64(n)
	wc.PrintProgress()
	return n, nil
}

func (wc WriteCounter) PrintProgress() {
	// Clear the line by using a character return to go back to the start and remove
	// the remaining characters by filling it with spaces
	fmt.Printf("\r%s", strings.Repeat(" ", 35))

	// Return again and print current status of download
	// We use the humanize package to print the bytes in a meaningful way (e.g. 10 MB)
	fmt.Printf("\rDownloading... %s complete", humanize.Bytes(wc.Total))
}

// DownloadFile will download a url to a local file. It's efficient because it will
// write as it downloads and not load the whole file into memory. We pass an io.TeeReader
// into Copy() to report progress on the download.
func DownloadFile(filepath string, url string) error {

	// Create the file, but give it a tmp file extension, this means we won't overwrite a
	// file until it's downloaded, but we'll remove the tmp extension once downloaded.
	fmt.Println(filepath)
	out, err := os.Create(filepath + ".tmp")
	if err != nil {
		return err
	}

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		out.Close()
		return err
	}
	defer resp.Body.Close()

	// Create our progress reporter and pass it to be used alongside our writer
	counter := &WriteCounter{}
	if _, err = io.Copy(out, io.TeeReader(resp.Body, counter)); err != nil {
		out.Close()
		return err
	}

	// The progress use the same line so print a new line once it's finished downloading
	fmt.Print("\n")

	// Close the file without defer so it can happen before Rename()
	out.Close()

	if err = os.Rename(filepath+".tmp", filepath); err != nil {
		return err
	}

	zipFile, err := zip.OpenReader(filepath)
	if err != nil {
		log.Fatal(err)
	}
	defer zipFile.Close()

	for _, f := range zipFile.File {
		fmt.Printf("Contents of %s:\n", f.Name)
		rc, err := f.Open()
		if err != nil {
			log.Fatal(err)
		}
		_, err = io.CopyN(os.Stdout, rc, 68)
		if err != nil {
			log.Fatal(err)
		}
		rc.Close()
		fmt.Println()
	}

	return nil
}

func updateDisasterCli() {

	executeDownload, err := checkForNewRelease()
	if err != nil {
		fmt.Println(err)
	}

	if executeDownload {

		// Get the current location of where the binary is located
		binDir, err := os.Getwd()
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(os.TempDir())
		fmt.Println("Installing new version of disaster-cli at ", binDir)

		// Get the download URL for the correct release file
		downloadUrl, downloadFileName, err := getReleaseURL()
		if err != nil {
			fmt.Println(err)
		}

		// Download the proper release asset (zip)
		if err := DownloadFile(downloadFileName, downloadUrl); err != nil {
			fmt.Println(err)
		}
	} else {
		fmt.Println("No new version found")
	}

}
