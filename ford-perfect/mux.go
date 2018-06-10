package main

import (
	"net/http"

	"go.opencensus.io/plugin/ochttp"
)

func createHTTPHandler() http.Handler {
	mux := http.NewServeMux()
	mux.Handle("/ready", readinessProbeHandler())
	mux.Handle("/healthz", livenessProbeHandler())
	return mux
}

func createHTTPSHalder(token string) http.Handler {
	mux := http.NewServeMux()
	mux.Handle("/updatesHook"+token, updatesHookHandler())
	return &ochttp.Handler{Handler: mux}
}
