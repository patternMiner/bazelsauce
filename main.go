package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/patternMiner/bazelsauce/context"
	"github.com/patternMiner/bazelsauce/handlers"
	"path/filepath"
)
var (
	bugs_data, _ = filepath.Abs("data/bugs.csv")
	devices_data, _ = filepath.Abs("data/devices.csv")
	testers_data, _ = filepath.Abs("data/testers.csv")
	tester_device_data, _ = filepath.Abs("data/tester_device.csv")
	data_file_paths = []string {testers_data, devices_data, tester_device_data, bugs_data}
)

func main() {
	err := context.InitContext(data_file_paths)
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
