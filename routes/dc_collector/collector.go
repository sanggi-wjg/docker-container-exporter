package dc_collector

import (
	"github.com/prometheus/client_golang/prometheus"
)

type dcCollector struct {
	dockerStatusMetric *prometheus.Desc
}

func NewDcCollector() *dcCollector {
	return &dcCollector{
		dockerStatusMetric: prometheus.NewDesc("app_docker_container_status", "Shows status docker container", nil, nil),
	}
}

func (collector *dcCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- collector.dockerStatusMetric
}

func (collector *dcCollector) Collect(ch chan<- prometheus.Metric) {
	var metricValue float64
	ch <- prometheus.MustNewConstMetric(collector.dockerStatusMetric, prometheus.CounterValue, metricValue)
}

// var dockerStatusMetric = prometheus.NewGauge(prometheus.GaugeOpts{
// 	Name: "app_docker_container_status",
// 	Help: "Shows status docker container",
// })

// func Init() {
// prometheus.MustRegister(dockerStatusMetric)
// prometheus.MustRegister(
// collectors.NewProcessCollector(collectors.ProcessCollectorOpts{}),
// )
// dockerStatusMetric.Set(1)
// }
