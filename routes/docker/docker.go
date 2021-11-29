package docker

import (
	"context"
	"fmt"
	"log"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

// var cli *client.Client

// func Init() {
// 	var err error
// 	cli, err = client.NewClientWithOpts(client.FromEnv)
// 	if err != nil {
// 		panic(err)
// 	}
// }

// func GetDockerContainers() string {
// 	containers, err := cli.ContainerList(context.Background(), types.ContainerListOptions{})
// 	if err != nil {
// 		log.Fatalln(err.Error())
// 	}

// 	var result string
// 	for _, container := range containers {
// 		// fmt.Print(container.Image)
// 		result += fmt.Sprintf("%s %s %s %v %v\n\n", container.State, container.Status, container.Image, container.Ports, container.Mounts)
// 	}
// 	return result
// }

func GetDockerContainerStatus() {
	go func() {
		cli, err := client.NewClientWithOpts(client.FromEnv)
		if err != nil {
			panic(err)
		}
		containers, err := cli.ContainerList(context.Background(), types.ContainerListOptions{})
		if err != nil {
			log.Fatalln(err.Error())
		}
		for _, container := range containers {
			fmt.Printf("%s %s %s %v %v\n\n", container.State, container.Status, container.Image, container.Ports, container.Mounts)
		}
	}()
}
