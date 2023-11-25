package main

import (
	"log"
	"net/http"

	"go.opentelemetry.io/otel/exporters/prometheus"
	"go.opentelemetry.io/otel/metric/global"
	"go.opentelemetry.io/otel/metric/instrument"
)

func main() {
	exporter := prometheus.New()
	global.SetMeterProvider(exporter.MeterProvider())

	meter := global.Meter("test-meter")

	requestCounter, _ := meter.SyncInt64().Counter(
		"api_requests_total",
		instrument.WithDescription("Total number of requests"),
	)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		requestCounter.Add(r.Context(), 1)
		w.Write([]byte("Hello, World!"))
	})

	http.HandleFunc("/metrics", exporter.ServeHTTP)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
