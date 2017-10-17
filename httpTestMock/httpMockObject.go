package main

import (
	"net/http"
)

type Server interface {

	Request(string) (http.ResponseWriter, error)

}






