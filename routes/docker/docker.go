package docker

import (
	"context"
	"fmt"
	"log"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

func GetDockerContainers() []types.Container {
	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		log.Fatalln(err)
	}
	containers, err := cli.ContainerList(context.Background(), types.ContainerListOptions{All: true})
	if err != nil {
		log.Fatalln(err.Error())
	}
	return containers
}

func GetDockerContainerStatus() string {
	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		log.Fatalln(err)
	}
	containers, err := cli.ContainerList(context.Background(), types.ContainerListOptions{})
	if err != nil {
		log.Fatalln(err.Error())
	}

	var results string
	for _, container := range containers {
		results += fmt.Sprintf("%s %s %s %v %v %d\n\n",
			container.State, container.Status, container.Image, container.Ports, container.Mounts, container.Created)
	}
	return results
}
