package handlers

import (
	"net/http"
	"encoding/json"
	"fmt"
	"github.com/patternMiner/bazelsauce/context"
)

func CountriesHandler (w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	setAccessControlResponseHeaders(w, req)
	countries := context.StringSlice{}
	for country := range context.CountryList {
		countries = append(countries, country)
	}
	data, err := json.Marshal(Response{Items: countries})
	if err != nil {
		fmt.Fprintln(w, err)
	}
	fmt.Fprintln(w, string(data))
}
