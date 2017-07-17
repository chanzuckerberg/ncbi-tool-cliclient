package main

import (
	"log"
	"path/filepath"
	"net/url"
	"fmt"
	"os"
	"io"
	"github.com/urfave/cli"
	"net/http"
)

func fileSimple(c *cli.Context) {
	params := url.Values{}
	pathName := c.Args().First()
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

func fileHistory(c *cli.Context) {
	// Setup
	params := url.Values{}
	pathName := c.Args().First()
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

func downloadFromURL(url string, downloadName string) {
	// Download the file
	fmt.Println("Downloading " + downloadName + " ...")
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal("Error in HTTP get request. ", err)
	}

	// Create local file
	if dest != "" {
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