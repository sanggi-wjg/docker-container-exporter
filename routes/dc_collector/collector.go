package dc_collector

import (
	"strconv"
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
			[]string{"container_name"},
			nil,
		),
		containerUpTime: prometheus.NewDesc(
			"app_docker_container_up_time",
			"docker container uptime",
			[]string{"container_name"},
			nil,
		),
		containerCreateTime: prometheus.NewDesc(
			"app_docker_container_create_time",
			"docker container create time",
			[]string{"container_name"},
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
	containers := docker.GetDockerContainers()
	for _, container := range containers {
		containerInfo := newContainerInfo(container)
		// fmt.Println(containerInfo, container.State, container.Status)

		ch <- prometheus.MustNewConstMetric(collector.containerUp, prometheus.CounterValue, float64(containerInfo.IsUp), containerInfo.Name)
		ch <- prometheus.MustNewConstMetric(collector.containerUpTime, prometheus.CounterValue, float64(containerInfo.Uptime), containerInfo.Name)
		ch <- prometheus.MustNewConstMetric(collector.containerCreateTime, prometheus.CounterValue, containerInfo.CreatedTime, containerInfo.Name)
	}
}

type containerInfo struct {
	IsUp        int
	Uptime      int
	Name        string
	CreatedTime float64
}

func newContainerInfo(container types.Container) *containerInfo {
	isUp := parseToIsUp(container.State)
	uptime := parseToUptime(container.Status)
	name := strings.Replace(container.Names[0], "/", "", -1)
	createdTime := float64(container.Created)

	ci := containerInfo{
		IsUp:        isUp,
		Uptime:      uptime,
		Name:        name,
		CreatedTime: createdTime,
	}
	return &ci
}

func parseToUptime(status string) int {
	if strings.Contains(status, "Up") {
		st := strings.ReplaceAll(status, "Up", "")
		st = strings.ReplaceAll(st, " ", "")

		if strings.Contains(st, "minutes") {
			st = strings.ReplaceAll(st, "minutes", "")
			min, err := strconv.Atoi(st)
			if err != nil {
				return 1 * 60
			}
			return min * 60

		} else if strings.Contains(st, "hours") {
			st = strings.ReplaceAll(st, "hours", "")
			hour, _ := strconv.Atoi(st)
			return hour * 60 * 60

		} else if strings.Contains(st, "days") {
			st = strings.ReplaceAll(st, "days", "")
			day, _ := strconv.Atoi(st)
			return day * 60 * 60 * 24
		}
	}

	return 0
}

func parseToIsUp(state string) int {
	// "paused", "restarting", "running", "removing", "dead", "created", "exited"
	for _, s := range []string{"running", "restarting"} {
		if state == s {
			return 1
		}
	}
	return 0
}
