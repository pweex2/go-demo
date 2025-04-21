package main

import (
	"log"
	"net/http"

	"github.com/vvx2/go-demo/handler"
)

func main() {
	http.HandleFunc("/hello", handler.HelloHandler)

	log.Println("Starting server on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
