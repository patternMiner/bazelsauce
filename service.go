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

	http.HandleFunc("/info", handlers.InfoHandler)

	http.HandleFunc("/countries", handlers.CountriesHandler)

	http.HandleFunc("/devices", handlers.DevicesHandler)

	http.HandleFunc("/tester_match", handlers.MatchHandler)

	fs := http.FileServer(http.Dir("applause/client/tester-match/dist"))
	http.Handle("/applause/client/tester-match/dist/",
		http.StripPrefix("/applause/client/tester-match/dist/", fs))

	fmt.Println("Starting up the tester_match service on 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
