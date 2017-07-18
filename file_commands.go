package main

import (
	"fmt"
	"github.com/urfave/cli"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"
)

// fileSimple handles simple file info or download requests.
func fileSimple(c *cli.Context) {
	params := url.Values{}
	pathName := getPath(c)
	endpoint := server + "/file"
	if pathName != "" { // Required
		params.Set("path-name", pathName)
	} else {
		log.Fatal("No path name provided.")
	}
	if versionNum != "" {
		params.Set("version-num", versionNum)
	}
	fileRequest(endpoint, params)
}

// fileAtTime handles requests for a file at a point in time.
func fileAtTime(c *cli.Context) {
	params := url.Values{}
	pathName := c.Args().First()
	endpoint := server + "/file/at-time"
	if pathName != "" { // Required
		params.Set("path-name", pathName)
	} else {
		log.Fatal("No path name provided.")
	}
	if inputTime != "" { // Required
		params.Set("input-time", inputTime)
	} else {
		log.Fatal("No input time provided.")
	}
	fileRequest(endpoint, params)
}

// fileRequest provides file info or downloads a file based on the query.
func fileRequest(endpoint string, params url.Values) {
	req := endpoint + "?" + params.Encode()
	fmt.Println("Request: " + req)
	if !download {
		res := getRequestToJSON(req)
		fmt.Println("File Info:\n", res)
	} else {
		res := getRequestToEntry(req)
		name := filepath.Base(res.Path)
		downloadFromURL(res.URL, name)
	}
}

// fileHistory outputs the version history of a file.
func fileHistory(c *cli.Context) {
	// Setup
	params := url.Values{}
	pathName := getPath(c)
	endpoint := server + "/file/history"
	if pathName != "" { // Required
		params.Set("path-name", pathName)
	} else {
		log.Fatal("No path name provided.")
	}

	// Request
	res := paramsToRequest(endpoint, params)
	fmt.Println("File History:", res)
}

// downloadFromURL downloads a file from a direct download URL to local disk.
func downloadFromURL(url string, downloadName string) {
	// Download the file
	fmt.Println("Downloading " + downloadName + " ...")
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal("Error in HTTP get request. ", err)
	}

	// Create local file
	if dest != "" {
		dest = strings.TrimSuffix(dest, "/")
		// Ex: $HOME/Desktop/FASTA/env_nr.gz
		downloadName = dest + "/" + downloadName
	}
	file, err := os.Create(downloadName)
	if err != nil {
		log.Fatal("Error in making local file. ", err)
	}
	_, err = io.Copy(file, resp.Body)
	if err != nil {
		log.Fatal("Error in writing local file. ", err)
	}
	fmt.Println("File downloaded.")
	err = file.Close()
	if err != nil {
		log.Fatal("Error in closing local file. ", err)
	}
}
