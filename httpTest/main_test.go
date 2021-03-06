package httpTest

import (
	"testing"
	"net/http"
	"encoding/json"
	"io/ioutil"
	"bytes"
	"log"
	"fmt"
	"net/http/httptest"
	"time"
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

func StartHttpServer() *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK) // implicit
			w.Header().Add("content-type", "application/json")
			m := Message{"some content", "some more content"}
			content, err := json.Marshal(m)
			if err != nil {
				log.Fatal("Error when marshalling: " + err.Error())
			}
			w.Write(content)
	}))
}


func TestRequest(t *testing.T) {

	defer Benchmark(time.Now())

	m := Request(url, GetFromHere)

	// Start http server
	//server := StartHttpServer()
	// Get URL
	//url = server.URL
	//defer server.Close()

	//m := Request(url, http.DefaultClient.Get)

	if m.Content != "some content" {
		t.Errorf("Error when receiving data: %v", m.Content)
	}

}
