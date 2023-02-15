package main

import (
	"fmt"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
	"time"
)

var MetricsName = []string{"ConnectionUsage", "CpuUsage", "DiskUsage", "IOPSUsage", "MemoryUsage", "MySQL_ActiveSessions"}

type Metrics struct {
	Metrics   []*prometheus.Desc
	DataPoint []map[string]interface{}
}

// define to global variabel
var configfile ConfigFile

func main() {
	configfile.ConfigRead("./config.json")
	dp := Metrics{}
	dp.MetricData(configfile.Instance_id)
	// 使用goroutine更新m.datapoint数据
	go func() {
		for {
			time.Sleep(90 * time.Second)
			dp.MetricData(configfile.Instance_id)
		}
	}()

	// 用于检验m.datapoint中的数据是否成功更新，debug使用。正常运行时注释此代码
	//go func() {
	//	for {
	//		time.Sleep(30 * time.Second)
	//		fmt.Println("listening datapoint goroutine.....")
	//		fmt.Println(dp.DataPoint)
	//	}
	//}()
	registry := prometheus.NewRegistry()
	registry.Register(dp.NewMetirc())
	//fmt.Println("the dp.Metrics data is ,", dp.Metrics)
	http.Handle("/metrics", promhttp.HandlerFor(registry, promhttp.HandlerOpts{Registry: registry}))
	http.ListenAndServe(":9999", nil)
}

func (m *Metrics) NewMetirc() prometheus.Collector {
	var desc []*prometheus.Desc
	for _, value := range m.DataPoint {
		name := fmt.Sprintf("%s", value["metrics_name"])
		id := fmt.Sprintf("%s", value["instanceId"])
		desc = append(desc, prometheus.NewDesc(name, "help", nil, prometheus.Labels{"instance": id, "metrics_name": name}))
	}
	m.Metrics = desc
	//fmt.Println("the desc[] is", desc)
	return m
}

func (m *Metrics) Describe(ch chan<- *prometheus.Desc) {
	for _, metric := range m.Metrics {
		ch <- metric
	}
}

func (m *Metrics) Collect(ch chan<- prometheus.Metric) {
	for index, metric := range m.Metrics {
		value := m.DataPoint[index]["Average"].(float64)
		ch <- prometheus.MustNewConstMetric(metric, prometheus.GaugeValue, value)

	}
}
