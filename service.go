package main

import (
	"./context"
	"fmt"
	"./handlers"
	"log"
	"net/http"
)

func main() {
	err := context.InitContext()
	if err != nil {
		fmt.Println(err)
	}

	http.HandleFunc("/", handlers.DefaultHandler)

	http.HandleFunc("/info", handlers.InfoHandler)

	http.HandleFunc("/match", handlers.MatchHandler)

	fmt.Println("Starting up on 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
