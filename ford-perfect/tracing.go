package main

import (
	"os"

	"go.opencensus.io/exporter/stackdriver"
	"go.opencensus.io/stats/view"
	"go.opencensus.io/trace"
)

func setupOpenCensus() error {
	// Create exporter
	se, err := stackdriver.NewExporter(stackdriver.Options{
		ProjectID: os.Getenv("GOOGLE_CLOUD_PROJECT"),
	})
	if err != nil {
		return err
	}
	trace.RegisterExporter(se)
	view.RegisterExporter(se)

	// Set the tracer
	trace.ApplyConfig(trace.Config{DefaultSampler: trace.AlwaysSample()})

	return nil
}
