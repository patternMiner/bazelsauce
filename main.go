package main

import (
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"
	"github.com/patternMiner/bazelsauce/context"
	"github.com/patternMiner/bazelsauce/handlers"
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

	fs := http.FileServer(http.Dir("client/tester-match/dist"))
	http.Handle("/static/",
		http.StripPrefix("/static/", fs))

	fmt.Println("Starting up the tester_match service on 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
