package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	log.Println("Server started on port 9090")
	err := http.ListenAndServe(":9090", http.FileServer(http.Dir("../../assets")))
	if err != nil {
		fmt.Println("Failed to start server", err)
		return
	}
}
