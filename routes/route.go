package routes

import (
	"fmt"
	"net/http"

	"github.com/docker/docker/client"
)

func GetMetrics(w http.ResponseWriter, req *http.Request) {
	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		panic(err)
	}
	fmt.Println(cli)
	w.Write([]byte("Hello World"))
}
