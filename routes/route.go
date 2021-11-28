package routes

import (
	"context"
	"fmt"
	"net/http"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

func Ping(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("pong"))
}

func GetContainers(w http.ResponseWriter, req *http.Request) {
	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		panic(err)
	}

	containers, err := cli.ContainerList(context.Background(), types.ContainerListOptions{})
	for _, container := range containers {
		fmt.Println(container)
		fmt.Printf("%s %s\n", container.ID[:10], container.Image)
	}
}
