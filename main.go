package main

import (
	"log"
	"net/http"
	"server/handler"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/check", handler.HealthCheck)
	s := &http.Server{
		Addr:    ":9777",
		Handler: mux,
	}
	log.Fatal(s.ListenAndServe())
}
