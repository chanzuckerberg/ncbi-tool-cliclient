package main

import (
	"github.com/urfave/cli"
	"net/url"
	"log"
	"fmt"
	"path/filepath"
	"encoding/json"
	"os"
)

func directorySimple(c *cli.Context) {
	// Setup
	params := url.Values{}
	pathName := c.Args().First()
	endpoint := server + "/directory"
	if pathName != "" { // Required
		params.Set("path-name", pathName)
	} else {
		log.Fatal("No path name provided.")
	}

	// Request
	if !download {
		res := paramsToRequest(endpoint, params)
		fmt.Println("Directory Info:", res)
	} else {
		// Get listing from server
		params.Set("output", "with-URLs")
		req := endpoint + "?" + params.Encode()
		fmt.Println("Request: " + req)
		fmt.Println("Downloading directory...")
		body := getRequestToBody(req)
		var listing []map[string]interface{}
		err := json.Unmarshal([]byte(body), &listing)
		if err != nil {
			log.Fatal("Error in unmarshalling response. ", err)
		}

		// Make folder
		dir := filepath.Base(pathName) // Ex: FASTA
		var sub string
		if dest != "" {
			sub = dest + "/" + dir // Ex: $HOME/Desktop/FASTA
		}
		fmt.Println("Making sub-folder " + sub + "/ ...")
		os.Mkdir(sub, os.ModePerm)

		// Download file-by-file
		for _, entry := range listing {
			path := fmt.Sprintf("%s", entry["Path"])
			url := fmt.Sprintf("%s", entry["URL"])
			name := filepath.Base(path)
			downloadFromURL(url, dir + "/" + name)
		}
		fmt.Println("Done downloading all files.")
	}
}

func directoryAtTime(c *cli.Context) {
	// Setup
	params := url.Values{}
	pathName := c.Args().First()
	endpoint := server + "/directory/at-time"
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

	// Request
	if !download {
		res := paramsToRequest(endpoint, params)
		fmt.Println("Directory Info:", res)
	} else {
		// Get listing from server
		params.Set("output", "with-URLs")
		req := endpoint + "?" + params.Encode()
		fmt.Println("Request: " + req)
		fmt.Println("Downloading directory...")
		body := getRequestToBody(req)
		var listing []map[string]interface{}
		err := json.Unmarshal([]byte(body), &listing)
		if err != nil {
			log.Fatal("Error in unmarshalling response. ", err)
		}

		// Make folder
		dir := filepath.Base(pathName) // Ex: FASTA
		var sub string
		if dest != "" {
			sub = dest + "/" + dir // Ex: $HOME/Desktop/FASTA
		}
		fmt.Println("Making sub-folder " + sub + "/ ...")
		err = os.MkdirAll(sub, os.ModePerm)
		if err != nil {
			log.Fatal("Error in making sub-folder. ", err)
		}

		// Download file-by-file
		for _, entry := range listing {
			path := fmt.Sprintf("%s", entry["Path"])
			url := fmt.Sprintf("%s", entry["URL"])
			name := filepath.Base(path)
			downloadFromURL(url, dir + "/" + name)
		}
		fmt.Println("Done downloading all files.")
	}
}

func directoryCompare(c *cli.Context) {
	// Setup
	params := url.Values{}
	pathName := c.Args().First()
	endpoint := server + "/directory/compare"
	if pathName != "" { // Required
		params.Set("path-name", pathName)
	} else {
		log.Fatal("No path name provided.")
	}
	if startDate != "" { // Required
		params.Set("start-date", startDate)
	} else {
		log.Fatal("No start datetime provided.")
	}
	if endDate != "" { // Required
		params.Set("end-date", endDate)
	} else {
		log.Fatal("No end datetime provided.")
	}

	// Request
	res := paramsToRequest(endpoint, params)
	fmt.Println("Directory Info:", res)
}