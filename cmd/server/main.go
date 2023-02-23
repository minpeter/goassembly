package main

import (
	"log"
	"net/http"
)

func main() {
	fs := http.FileServer(http.Dir("./docs"))
	http.Handle("/", fs)

	log.Println("Listening on http://localhost:9090/index.html")
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal(err)
	}
}
