package main

import (
	"log"
	"net/http"
)

func main() {

	mux := http.NewServeMux()
	mux.Handle("/ready", http.HandlerFunc(readinessProbeHandler))
	mux.Handle("/healthz", http.HandlerFunc(livenessProbeHandler))

	log.Fatal(http.ListenAndServe(":8080", mux))
}
