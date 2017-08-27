package handlers

import (
	"net/http"
	"encoding/json"
	"fmt"
	"github.com/patternMiner/bazelsauce/context"
)

func MatchHandler (w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	setAccessControlResponseHeaders(w, req)
	query := req.URL.Query()
	countries := query["country"]
	devices := query["device"]
	testersByRank := context.MatchTesters(countries, devices)
	testers := make([]Tester, len(testersByRank))
	for i, pair := range testersByRank {
		tester := context.TesterMap[pair.Key]
		rank := pair.Value
		testers[i] = Tester{Id: tester[0], FirstName: tester[1], LastName: tester[2], Country: tester[3], Rank: rank}
	}
	data, err := json.Marshal(Response{Items: testers})
	if err != nil {
		fmt.Fprintln(w, err)
	}
	fmt.Fprintln(w, string(data))
}
