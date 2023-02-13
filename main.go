package main

import (
	"github.com/gophermasters/x/handlers"
	"net/http"
)

const PORT = os.Getenv("PORT")

func main(){
	http.HandleFunc("/", handlers.ServeHTML)
	http.ListenAndServe(PORT, nil)
}