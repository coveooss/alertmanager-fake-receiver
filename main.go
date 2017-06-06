package main

import (
	"fmt"
	"log"
	"net/http"
)

func fakeResponse(w http.ResponseWriter, r *http.Request) {
	log.Print("Alert reveived")
	fmt.Fprintf(w, "Fake Alert")
}

func main() {
	log.Print("Starting fake service")

	http.HandleFunc("/", fakeResponse)
	err := http.ListenAndServe(":80", nil)

	if err != nil {
		log.Fatal("Cannot start service:", err)
	}
}
