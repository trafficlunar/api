package server

import (
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func StartPrometheusServer() {
	go func() {
		http.HandleFunc("/", promhttp.Handler().ServeHTTP)
		http.ListenAndServe(":8889", nil)
	}()
}
