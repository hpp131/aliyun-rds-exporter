package main

import (
	"fmt"
	"github.com/prometheus/client_golang/prometheus"
)

const (
	AccessKey = "LTAI5tPwoRHDHjZDCY85QPN8"
	SecretKey = "WBEOSMwaqljhBQlqoi5WfTFqVe45wg"
)

var MetricsName = []string{"ConnectionUsage", "CpuUsage", "DiskUsage", "IOPSUsage", "MemoryUsage", "MySQL_ActiveSessions"}

type AllData struct {
	DataPoint []map[string]interface{}
}

type Metrics struct {
	Metrics []*prometheus.Desc
}

func main() {
	dp := AllData{}
	//point := &dp
	dp.MetricData(MetricsName)
	fmt.Println(dp.DataPoint)

}

func (dp *AllData) NewMetirc() *Metrics {
	var desc = []*prometheus.Desc{}
	//var variablelabel=  []string{"instanceID"}
	var constantlabel = map[string]string{"instanceid": "rm-uf64if27mi7e9p63l"}
	for _, value := range dp.DataPoint {
		name := fmt.Sprintf("%s", value["MetircName"])
		desc = append(desc, prometheus.NewDesc(name, "help", nil, constantlabel))
	}
	return &Metrics{
		Metrics: desc,
	}
}

func (m *Metrics) Describe(ch <-chan *prometheus.Desc) {
	for _, metric := range m.Metrics {
		ch <- metric
	}
}

func (m *Metrics) Collect(ch <-chan prometheus.Metric, dp *AllData) {
	for index, metric := range m.Metrics {
		ch <- prometheus.MustNewConstMetric(metric, prometheus.GaugeValue, dp.DataPoint[index][], nil)

	}
}
