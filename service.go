package main

import (
	"github.com/patternMiner/applause/context"
	"fmt"
	"github.com/patternMiner/applause/handlers"
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

	fs := http.FileServer(http.Dir("github.com/patternMiner/applause/client/tester-match/dist"))
	http.Handle("/static/",
		http.StripPrefix("/static/", fs))

	fmt.Println("Starting up the tester_match service on 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
