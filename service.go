package main

import (
	"./context"
	"fmt"
	"log"
	"net/http"
)

func main() {
	err := context.InitContext()
	if err != nil {
		fmt.Println(err)
	}

	http.HandleFunc("/", defaultHandler)

	http.HandleFunc("/info", infoHandler)

	http.HandleFunc("/match", matchHandler)

	fmt.Println("Starting up on 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func defaultHandler (w http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(w, "Hello World from docker!")
}

func matchHandler (w http.ResponseWriter, req *http.Request) {
	query := req.URL.Query()
	countries := query["country"]
	devices := query["device"]
	testers := context.TestersByCountryDevice(countries, devices)
	testerRankMap := context.RankTesters(testers, devices)
	testersByRank := context.SortByRank(testerRankMap)
	fmt.Fprintf(w, "Testers by rank: %v\n", testersByRank)
}

func infoHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(w, "Testers")
	for id, tester := range context.TesterMap {
		fmt.Fprintf(w, "%2s: %s\n", id, tester)
	}
	fmt.Fprintln(w)

	fmt.Fprintln(w, "Testers by country")
	for country, testers := range context.CountryTestersMap {
		fmt.Fprintf(w, "%s: %s\n", country, testers)
	}
	fmt.Fprintln(w)

	fmt.Fprintln(w, "Devices")
	for id, device := range context.DeviceMap {
		fmt.Fprintf(w, "%2s: %s\n", id, device)
	}
	fmt.Fprintln(w)

	fmt.Fprintln(w, "Testers by device")
	for device, testers := range context.DeviceTestersMap {
		fmt.Fprintf(w, "%2s: %s\n", device, testers)
	}
	fmt.Fprintln(w)

	fmt.Fprintln(w, "Testers by country_device")
	for country_device, testers := range context.CountryDeviceTestersMap {
		fmt.Fprintf(w, "%8s: %s\n", country_device, testers)
	}
	fmt.Fprintln(w)

	fmt.Fprintln(w, "BugCount by tester_device")
	for tester_device, count := range context.TesterDeviceBugCountMap {
		fmt.Fprintf(w, "%5s: %d\n", tester_device, count)
	}
	fmt.Fprintln(w)
}
