package routes

import (
	"context"
	"fmt"
	"net/http"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"

	log "github.com/sirupsen/logrus"
)

func Ping(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("pong"))
}

func GetContainers(w http.ResponseWriter, req *http.Request) {
	defer log.Debug("route-GetContainers")

	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		panic(err)
	}

	containers, err := cli.ContainerList(context.Background(), types.ContainerListOptions{})
	if err != nil {
		log.Fatal(err.Error())
	}

	var result string
	for _, container := range containers {
		// fmt.Print(container.Image)
		result += fmt.Sprintf("%s %s %s %v %v\n", container.State, container.Status, container.Image, container.Ports, container.Mounts)
	}
	w.Write([]byte(result))
}
