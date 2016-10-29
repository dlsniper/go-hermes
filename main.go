package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome to our API!"))
	})

	log.Fatalln(http.ListenAndServe(":80", nil))
}
