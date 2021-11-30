package dc_collector

import (
	"strings"

	"github.com/docker/docker/api/types"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/sanggi-wjg/docker-continer-exporter/routes/docker"
)

type containerMetric struct {
	containerUp         *prometheus.Desc
	containerUpTime     *prometheus.Desc
	containerCreateTime *prometheus.Desc
}

// var metricLabels []string{"names"}

func NewContainerMetric() *containerMetric {
	return &containerMetric{
		containerUp: prometheus.NewDesc(
			"app_docker_container_up",
			"docker container up",
			[]string{"names"},
			nil,
		),
		containerUpTime: prometheus.NewDesc(
			"app_docker_container_up_time",
			"docker container uptime",
			[]string{"names"},
			nil,
		),
		containerCreateTime: prometheus.NewDesc(
			"app_docker_container_create_time",
			"docker container create time",
			[]string{"names"},
			nil,
		),
	}
}

func (collector *containerMetric) Describe(ch chan<- *prometheus.Desc) {
	ch <- collector.containerUp
	ch <- collector.containerUpTime
	ch <- collector.containerCreateTime
}

func (collector *containerMetric) Collect(ch chan<- prometheus.Metric) {
	var metricValue float64

	containers := docker.GetDockerContainers()
	for _, container := range containers {
		containerInfo := newContainerInfo(container)
		// fmt.Println(containerInfo, container.State, container.Status)

		ch <- prometheus.MustNewConstMetric(collector.containerUp, prometheus.CounterValue, float64(containerInfo.IsUp), containerInfo.Name)
		ch <- prometheus.MustNewConstMetric(collector.containerUpTime, prometheus.CounterValue, metricValue, containerInfo.Name)
		ch <- prometheus.MustNewConstMetric(collector.containerCreateTime, prometheus.CounterValue, containerInfo.CreatedTime, containerInfo.Name)
	}
}

type containerInfo struct {
	IsUp        int
	Name        string
	CreatedTime float64
}

func newContainerInfo(container types.Container) *containerInfo {
	// fmt.Println(container.Status)
	isUp := parseStateToInt(container.State)
	name := strings.Replace(container.Names[0], "/", "", -1)
	createdTime := float64(container.Created)

	ci := containerInfo{IsUp: isUp, Name: name, CreatedTime: createdTime}
	return &ci
}

func parseStateToInt(state string) int {
	// "paused", "restarting", "running", "removing", "dead", "created", "exited"
	for _, s := range []string{"running", "restarting", "removing"} {
		if state == s {
			return 1
		}
	}
	return 0
}
