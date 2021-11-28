package main

import (
	"net/http"

	"github.com/sanggi-wjg/docker-continer-client/routes"
)

func main() {
	http.HandleFunc("/metrics", routes.GetMetrics)
	http.ListenAndServe(":9091", nil)
}
