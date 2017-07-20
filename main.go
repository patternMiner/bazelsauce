package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/patternMiner/applause/service"
)

func main() {
	err := service.InitContext()
	if err != nil {
		fmt.Println(err)
	}

	http.HandleFunc("/info", service.InfoHandler)

	http.HandleFunc("/countries", service.CountriesHandler)

	http.HandleFunc("/devices", service.DevicesHandler)

	http.HandleFunc("/tester_match", service.MatchHandler)

	fs := http.FileServer(http.Dir("github.com/patternMiner/applause/client/tester-match/dist"))
	http.Handle("/static/",
		http.StripPrefix("/static/", fs))

	fmt.Println("Starting up the tester_match service on 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
