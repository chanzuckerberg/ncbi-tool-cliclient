package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

// getRequestToEntry takes in a GET request URL and returns an Entry.
func getRequestToEntry(input string) Entry {
	res := Entry{}
	body := getRequestToBody(input)
	err := json.Unmarshal(body, &res)
	if err != nil {
		log.Fatal("Error in marshalling entry. ", err)
	}
	if res.Path == "" {
		log.Fatal("No results found.")
	}
	return res
}

// getRequestToJSON takes in a GET request URL and returns the output in
// formatted JSON.
func getRequestToJSON(input string) string {
	res := getRequestToBody(input)
	var pretty bytes.Buffer
	err := json.Indent(&pretty, res, "", "    ")
	if err != nil {
		log.Fatal("JSON parse error. ", err)
	}
	return string(pretty.Bytes())
}

// getRequestToBody takes in a GET request URL and returns the HTTP body in
// bytes.
func getRequestToBody(input string) []byte {
	var body []byte
	res, err := http.Get(input)
	if err != nil {
		log.Fatal("Error in HTTP get request. ", err)
	}
	defer func() {
		closeErr := res.Body.Close()
		if closeErr != nil {
			log.Fatal("Couldn't close HTTP response body. ", closeErr)
		}
	}()
	body, err = ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal("Error in reading HTTP body.")
	}
	return body
}

// paramsToRequest takes in the request values and returns the formatted JSON
// response.
func paramsToRequest(endpoint string, params url.Values) string {
	req := endpoint + "?" + params.Encode()
	fmt.Println("Request: " + req)
	return getRequestToJSON(req)
}
