package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/metrics", func(w http.ResponseWriter, req *http.Request) {

	})
	http.ListenAndServe(":9091", nil)
}
