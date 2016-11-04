package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/ggirtsou/go-hermes/model"
)

var payload model.Payload

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome to our API!"))

		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&payload)
		if err != nil {
			log.Fatalln(err)
		}

		// TODO(ggirtsou): process payload
		fmt.Println(payload)

		defer r.Body.Close()
	})

	log.Fatalln(http.ListenAndServe(":80", nil))
}
