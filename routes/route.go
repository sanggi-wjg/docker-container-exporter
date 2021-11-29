package routes

import (
	"net/http"

	"github.com/sanggi-wjg/docker-continer-exporter/routes/docker"
	log "github.com/sirupsen/logrus"
)

func Ping(w http.ResponseWriter, req *http.Request) {
	defer log.Debug("pong")
	w.Write([]byte("pong"))
}

func GetContainers(w http.ResponseWriter, req *http.Request) {
	results := docker.GetDockerContainerStatus()
	w.Write([]byte(results))
}
