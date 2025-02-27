package main

import (
	controllers2 "github.com/svondras/url-shortener-app.git/internal/controllers"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", controllers2.ShowIndex)
	http.HandleFunc("/shorten", controllers2.Shorten)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
