package main

import (
	"log"
	"net/http"

	"go.opentelemetry.io/otel/exporters/prometheus"
)

func main() {
	exporter, err := prometheus.New()
	if err != nil {
		log.Fatalf("failed to initialize prometheus exporter: %v", err)
	}

	// Cria um servidor HTTP para expor as métricas
	http.HandleFunc("/", handler)
	http.HandleFunc("/metrics", exporter.ServeHTTP)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	// Sua lógica de manipulação de requisições aqui
}
