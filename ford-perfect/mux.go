package main

import "net/http"

func createHTTPHandler() http.Handler {
	mux := http.NewServeMux()
	mux.Handle("/ready", readinessProbeHandler())
	mux.Handle("/healthz", livenessProbeHandler())
	return mux
}

func createHTTPSHalder(token string) http.Handler {
	mux := http.NewServeMux()
	mux.Handle("/updatesHook"+token, updatesHookHandler())
	return mux
}
