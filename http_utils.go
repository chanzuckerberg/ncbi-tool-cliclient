package main

import (
	"log"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"bytes"
	"fmt"
	"net/url"
)

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

func getRequestToJSON(input string) string {
	res := getRequestToBody(input)
	var pretty bytes.Buffer
	err := json.Indent(&pretty, res, "", "    ")
	if err != nil {
		log.Fatal("JSON parse error. ", err)
	}
	return string(pretty.Bytes())
}

func getRequestToBody(input string) []byte {
	// Get request
	var body []byte
	res, err := http.Get(input)
	if err != nil {
		log.Fatal("Error in HTTP get request. ", err)
	}
	defer func() {
		closeErr := res.Body.Close()
		if closeErr != nil {
			err = ComboErr("Couldn't close HTTP response body.", closeErr, err)
			log.Fatal(err)
		}
	}()
	body, err = ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal("Error in reading HTTP body.")
	}
	return body
}

func paramsToRequest(endpoint string, params url.Values) string {
	req := endpoint + "?" + params.Encode()
	fmt.Println("Request: " + req)
	return getRequestToJSON(req)
}