package dc_collector

import (
	"strings"

	"github.com/docker/docker/api/types"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/sanggi-wjg/docker-continer-exporter/routes/docker"
)

type containerMetric struct {
	containerIsUp      *prometheus.Desc
	containerCreatedAt *prometheus.Desc
}

// var metricLabels []string{"names"}

func NewContainerMetric() *containerMetric {
	return &containerMetric{
		containerIsUp: prometheus.NewDesc(
			"docker_container_is_up",
			"docker container is up",
			[]string{"container_id", "container_name", "image_name"},
			nil,
		),
		containerCreatedAt: prometheus.NewDesc(
			"docker_container_created_at",
			"docker container created time",
			[]string{"container_id", "container_name", "image_name"},
			nil,
		),
	}
}

func (collector *containerMetric) Describe(ch chan<- *prometheus.Desc) {
	ch <- collector.containerIsUp
	ch <- collector.containerCreatedAt
}

func (collector *containerMetric) Collect(ch chan<- prometheus.Metric) {
	for _, container := range docker.GetDockerContainers() {
		containerInfo := newDockerContainerInfo(container)

		ch <- prometheus.MustNewConstMetric(
			collector.containerIsUp, prometheus.CounterValue, float64(containerInfo.isUp),
			containerInfo.containerId, containerInfo.containerName, containerInfo.imageName,
		)
		ch <- prometheus.MustNewConstMetric(
			collector.containerCreatedAt, prometheus.CounterValue, containerInfo.createdAt,
			containerInfo.containerId, containerInfo.containerName, containerInfo.imageName,
		)
	}
}

type dockerContainerInfo struct {
	containerId   string
	containerName string
	imageName     string
	isUp          int
	createdAt     float64
}

func newDockerContainerInfo(container types.Container) *dockerContainerInfo {
	d := dockerContainerInfo{
		containerId:   container.ID,
		containerName: strings.Replace(container.Names[0], "/", "", -1),
		imageName:     container.Image,
		isUp:          parseToIsUp(container.State),
		createdAt:     float64(container.Created),
	}
	return &d
}

func parseToIsUp(state string) int {
	// "paused", "restarting", "running", "removing", "dead", "created", "exited"
	for _, s := range []string{"running"} {
		if state == s {
			return 1
		}
	}
	return 0
}
