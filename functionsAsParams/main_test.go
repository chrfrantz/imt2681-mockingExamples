package main

import (
	"testing"
	"net/http"
	"encoding/json"
	"io/ioutil"
	"bytes"
	"log"
	"fmt"
	"github.com/stretchr/testify/assert"
)

func GetFromHere(url string) (*http.Response, error) {
	resp := http.Response{}
	resp.StatusCode = http.StatusOK
	resp.Status = "whatever"
	resp.Header = http.Header{}
	resp.Header.Add("content-type", "application/json")
	m := Message{"some content", "some more content"}
	content, err := json.Marshal(m)
	if err != nil {
		log.Fatal("Error when marshalling: " + err.Error())
	} else {
		fmt.Println(string(content))
	}
	resp.Body = ioutil.NopCloser(bytes.NewBufferString(string(content)))
	return &resp, nil
}


func TestRequest(t *testing.T) {

	m := Request(url, GetFromHere)

	/// Using assert statements
	assert.NotNil(t, m, "M is nil")
	assert.Equal(t, "some content", m.Content, "Web service response does not correspond to 'some content'")

	/// Conventional testing
	//if m.Content != "some content" {
	//	t.Errorf("Error when receiving data: %v", m.Content)
	//}

}
