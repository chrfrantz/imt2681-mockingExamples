package main

import (
	"net/http/httptest"
	"net/http"
	"testing"
	"github.com/golang/mock/gomock"
	"io/ioutil"
	"bytes"
)

func TestHttpMock (t *testing.T) {
	mockctrl := gomock.NewController(t)
	defer mockctrl.Finish()

	mock := NewMockServer(mockctrl)

	mock.EXPECT().Request("some content").Times(1)

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
			content, err := ioutil.ReadAll(r.Body)
			if err != nil {
				t.Error("Error during decoding")
			}
			mock.Request(string(content))
	}))

	defer ts.Close()


	http.DefaultClient.Post(ts.URL, "text", ioutil.NopCloser(bytes.NewBufferString("some content")))

}
