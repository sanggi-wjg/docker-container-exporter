package dc_collector

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/sanggi-wjg/docker-continer-exporter/routes/docker"
)

type containerMetric struct {
	containerUp         *prometheus.Desc
	containerUpTime     *prometheus.Desc
	containerCreateTime *prometheus.Desc
	// containerStatus    *prometheus.Desc
	// containerImageName *prometheus.Desc
	// containerPorts     *prometheus.Desc
	// containerMounts    *prometheus.Desc
}

func NewContainerMetric() *containerMetric {
	return &containerMetric{
		containerUp: prometheus.NewDesc(
			"app_docker_container_up",
			"docker container up",
			[]string{"imageName"},
			nil,
		),
		containerUpTime: prometheus.NewDesc(
			"app_docker_container_up_time",
			"docker container uptime",
			[]string{"imageName"},
			nil,
		),
		containerCreateTime: prometheus.NewDesc(
			"app_docker_container_create_time",
			"docker container create time",
			[]string{"imageName"},
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
	// var names []string
	// for _, container := range containers {
	// 	names = append(names, container.Image)
	// }
	// if containers[0].Created

	ch <- prometheus.MustNewConstMetric(collector.containerUp, prometheus.CounterValue, metricValue, containers[0].Image)
	ch <- prometheus.MustNewConstMetric(collector.containerUpTime, prometheus.CounterValue, metricValue, containers[0].Image)
	ch <- prometheus.MustNewConstMetric(collector.containerCreateTime, prometheus.CounterValue, float64(containers[0].Created), containers[0].Image)
}
