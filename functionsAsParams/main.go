package main

import (
	"net/http"
	"log"
	"encoding/json"
	"io/ioutil"
	"fmt"
)

var url = "http://localhost:8080"

type Message struct {
	Content string `json:"content"`
	FurtherContent string `json:"furtherContent"`
}

func Request(url string, target func(string)(*http.Response, error)) Message {

	res, err := target(url)
	if err != nil {
		log.Fatal("Didn't work: " + err.Error())
	}

	body, err := ioutil.ReadAll(res.Body)

	m := Message{}
	json.Unmarshal(body, &m)

	fmt.Println(m)

	return m
}

func main() {

	Request(url, http.DefaultClient.Get)

}
