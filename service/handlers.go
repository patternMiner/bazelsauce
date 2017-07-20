package service

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Response struct {
	Items interface{}
	Err error
}

type Tester struct {
	Id string
	FirstName string
	LastName string
	Country string
	Rank int
}

type Device struct {
	Id string
	Description string
}

func DefaultHandler (w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	fmt.Fprintln(w, "/info\n\tShow context\n")
	fmt.Fprintln(w, "/tester_match?country=GB&device=5\n\tMatch testers for the country(GB) and device(5)")
}

func CountriesHandler (w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	countries := StringSlice{}
	for country := range CountryList {
		countries = append(countries, country)
	}
	data, err := json.Marshal(Response{Items: countries})
	if err != nil {
		fmt.Fprintln(w, err)
	}
	fmt.Fprintln(w, string(data))
}

func DevicesHandler (w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	devices := []Device{}
	for device := range DeviceList {
		deviceRecord := DeviceMap[device]
		devices = append(devices, Device{Id: deviceRecord[0], Description: deviceRecord[1]})
	}
	data, err := json.Marshal(Response{Items: devices})
	if err != nil {
		fmt.Fprintln(w, err)
	}
	fmt.Fprintln(w, string(data))
}

func MatchHandler (w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	query := req.URL.Query()
	countries := query["country"]
	devices := query["device"]
	testersByRank := MatchTesters(countries, devices)
	testers := make([]Tester, len(testersByRank))
	for i, pair := range testersByRank {
		tester := TesterMap[pair.Key]
		rank := pair.Value
		testers[i] = Tester{Id: tester[0], FirstName: tester[1], LastName: tester[2], Country: tester[3], Rank: rank}
	}
	data, err := json.Marshal(Response{Items: testers})
	if err != nil {
		fmt.Fprintln(w, err)
	}
	fmt.Fprintln(w, string(data))
}

func InfoHandler(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	fmt.Fprintln(w, "Testers")
	for id, tester := range TesterMap {
		fmt.Fprintf(w, "%2s: %s\n", id, tester)
	}
	fmt.Fprintln(w)

	fmt.Fprintln(w, "Testers by country")
	for country, testers := range CountryTestersMap {
		fmt.Fprintf(w, "%s: %s\n", country, testers)
	}
	fmt.Fprintln(w)

	fmt.Fprintln(w, "Devices")
	for id, device := range DeviceMap {
		fmt.Fprintf(w, "%2s: %s\n", id, device)
	}
	fmt.Fprintln(w)

	fmt.Fprintln(w, "Testers by device")
	for device, testers := range DeviceTestersMap {
		fmt.Fprintf(w, "%2s: %s\n", device, testers)
	}
	fmt.Fprintln(w)

	fmt.Fprintln(w, "Testers by country_device")
	for country_device, testers := range CountryDeviceTestersMap {
		fmt.Fprintf(w, "%8s: %s\n", country_device, testers)
	}
	fmt.Fprintln(w)

	fmt.Fprintln(w, "BugCount by tester_device")
	for tester_device, count := range TesterDeviceBugCountMap {
		fmt.Fprintf(w, "%5s: %d\n", tester_device, count)
	}
	fmt.Fprintln(w)
}
