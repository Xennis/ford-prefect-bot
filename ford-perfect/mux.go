package main

import "net/http"

func createHandler() http.Handler {
	mux := http.NewServeMux()
	mux.Handle("/ready", http.HandlerFunc(readinessProbeHandler))
	mux.Handle("/healthz", http.HandlerFunc(livenessProbeHandler))
	// TODO: append bot.Token
	mux.Handle("/updatesHook", http.HandlerFunc(updatesHookHandler))
	return mux
}
