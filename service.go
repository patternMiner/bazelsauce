package main

import (
	"./context"
	"fmt"
	"./handlers"
	"log"
	"net/http"
)

func main() {
	err := context.InitContext()
	if err != nil {
		fmt.Println(err)
	}

	http.HandleFunc("/", handlers.DefaultHandler)

	http.HandleFunc("/info", handlers.InfoHandler)

	http.HandleFunc("/countries", handlers.CountriesHandler)

	http.HandleFunc("/devices", handlers.DevicesHandler)

	http.HandleFunc("/tester_match", handlers.MatchHandler)

	fmt.Println("Starting up the tester_match service on 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
