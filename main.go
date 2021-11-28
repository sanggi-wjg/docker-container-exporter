package main

import (
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/sanggi-wjg/docker-continer-client/routes"
)

func main() {
	http.HandleFunc("/ping", routes.Ping)
	http.HandleFunc("/cs", routes.GetContainers)
	http.Handle("/metrics", promhttp.Handler())

	err := http.ListenAndServe(":9091", nil)
	if err != nil {
		panic(err)
	}
}
