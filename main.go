package main

import (
	"fmt"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
	"time"
)

const (
	AccessKey = "LTAI5tPwoRHDHjZDCY85QPN8"
	SecretKey = "WBEOSMwaqljhBQlqoi5WfTFqVe45wg"
)

var MetricsName = []string{"ConnectionUsage", "CpuUsage", "DiskUsage", "IOPSUsage", "MemoryUsage", "MySQL_ActiveSessions"}

type Metrics struct {
	Metrics   []*prometheus.Desc
	DataPoint []map[string]interface{}
}

func main() {
	var configfile ConfigFile
	configfile.ConfigRead("config.json")
	dp := Metrics{}
	dp.MetricData(configfile.Instance_id)
	// 使用goroutine更新m.datapoint数据
	go func() {
		for {
			time.Sleep(90 * time.Second)
			dp.MetricData(MetricsName)
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
	var constantlabel = map[string]string{"instanceid": "rm-uf64if27mi7e9p63l"}
	for _, value := range m.DataPoint {
		name := fmt.Sprintf("%s", value["MetircName"])
		desc = append(desc, prometheus.NewDesc(name, "help", nil, constantlabel))
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
