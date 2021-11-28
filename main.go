package main

import (
	"net/http"

	"github.com/sanggi-wjg/docker-continer-client/routes"
)

func main() {
	http.HandleFunc("/ping", routes.Ping)
	http.HandleFunc("/metrics", routes.GetMetrics)

	err := http.ListenAndServe(":9091", nil)
	if err != nil {
		panic(err)
	}
}
