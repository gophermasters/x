package main

import (
	"github.com/gophermasters/x/handlers"
	"net/http"
)


func main(){
	http.HandleFunc("/", handlers.ServeHTML)
	http.ListenAndServe(":10000", nil)
}