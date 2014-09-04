package main

import (
	"log"
	"net/http"
)

func main() {
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/", fs)

	log.Println("Listening on :8080...")
	http.ListenAndServe(":8080", nil)
}
