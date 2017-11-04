package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"
	"github.com/patternMiner/bazelsauce/context"
	"github.com/patternMiner/bazelsauce/handlers"
)

var (
	certPath string
	keyPath string
)

func init() {
	flag.StringVar(&certPath, "cert_path", "data/cert.pem", "SSL Certificate Path")
	flag.StringVar(&keyPath, "key_path", "data/key.pem", "SSL Key Path")
}

func main() {
	flag.Parse()

	err := context.InitContext()
	if err != nil {
		fmt.Println(err)
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		w.Header().Add("Strict-Transport-Security", "max-age=63072000; includeSubDomains")
		w.Write([]byte("This is an example server.\n"))
	})
	mux.HandleFunc("/info", handlers.InfoHandler)

	mux.HandleFunc("/countries", handlers.CountriesHandler)

	mux.HandleFunc("/devices", handlers.DevicesHandler)

	mux.HandleFunc("/tester_match", handlers.MatchHandler)

	fs := http.FileServer(http.Dir("client/tester-match/dist"))
	mux.Handle("/static/", http.StripPrefix("/static/", fs))

	fmt.Println("Starting up the tester_match http service on port 8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
