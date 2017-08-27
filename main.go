package main

import (
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"
	"github.com/patternMiner/bazelsauce/context"
	"github.com/patternMiner/bazelsauce/handlers"
	"crypto/tls"
)

func main() {
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

	cfg := &tls.Config{
		MinVersion:               tls.VersionTLS12,
		CurvePreferences:         []tls.CurveID{tls.CurveP521, tls.CurveP384, tls.CurveP256},
		PreferServerCipherSuites: true,
		CipherSuites: []uint16{
			tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
			tls.TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA,
			tls.TLS_RSA_WITH_AES_256_GCM_SHA384,
			tls.TLS_RSA_WITH_AES_256_CBC_SHA,
		},
	}
	srv := &http.Server{
		Addr:         ":8080",
		Handler:      mux,
		TLSConfig:    cfg,
		TLSNextProto: make(map[string]func(*http.Server, *tls.Conn, http.Handler), 0),
	}
	fmt.Println("Starting up the tester_match https service on port 8080")
	log.Fatal(srv.ListenAndServeTLS("localhost.pem", "server.key"))
}
