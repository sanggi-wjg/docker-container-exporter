package main

import (
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/prometheus/common/version"
	"github.com/sanggi-wjg/docker-continer-exporter/routes"
	"github.com/sanggi-wjg/docker-continer-exporter/routes/dc_collector"

	log "github.com/sirupsen/logrus"
)

func main() {
	prometheus.Register(version.NewCollector("docker_container_exporter"))
	if err := prometheus.Register(dc_collector.NewContainerMetric()); err != nil {
		log.Fatalln(err.Error())
	}

	http.HandleFunc("/ping", routes.Ping)
	http.HandleFunc("/containers", routes.GetContainers)
	http.Handle("/metrics", promhttp.Handler())

	log.Info("Starting http server with port:9091")
	log.Info("route to /ping /containers /metrics")

	err := http.ListenAndServe(":9091", nil)
	if err != nil {
		log.Fatal(err.Error())
	}
}
