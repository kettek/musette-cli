package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/url"
)

func requestPath(loc string, target interface{}) error {
	r, err := httpClient.Get(fmt.Sprintf("%s/api/browse%s", config.Server, url.QueryEscape(loc)))
	if err != nil {
		return err
	}
	defer r.Body.Close()

	return json.NewDecoder(r.Body).Decode(target)
}

func requestAPI(url string, target interface{}) error {
	r, err := httpClient.Get(url)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	return json.NewDecoder(r.Body).Decode(target)
}

func postAPI(url string, values url.Values, target interface{}) error {
	jsonBody, err := json.Marshal(values)
	r, err := httpClient.Post(url, "application/json", bytes.NewBuffer(jsonBody))

	if err != nil {
		panic(err)
	}
	defer r.Body.Close()

	/*if err != nil {
		return err
	}
	defer r.Body.Close()*/

	return json.NewDecoder(r.Body).Decode(target)
}
