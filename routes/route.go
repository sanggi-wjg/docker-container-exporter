package routes

import (
	"fmt"
	"net/http"

	"github.com/docker/docker/client"
)

func Ping(w http.ResponseWriter, req *http.Request) {
	fmt.Println("pong")
	w.Write([]byte("pong"))
}

func GetMetrics(w http.ResponseWriter, req *http.Request) {
	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		panic(err)
	}
	fmt.Println(cli)
	fmt.Println("hello worldddd")
	w.Write([]byte("Hello World"))
}
