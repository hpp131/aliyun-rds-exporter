package main

import "fmt"

const (
	AccessKey = "LTAI5tPwoRHDHjZDCY85QPN8"
	SecretKey = "WBEOSMwaqljhBQlqoi5WfTFqVe45wg"
)

var MetricsName = []string{"ConnectionUsage", "CpuUsage", "DiskUsage", "IOPSUsage", "MemoryUsage", "MySQL_ActiveSessions"}

type AllData struct {
	DataPoint []map[string]interface{}
}

func main() {
	dp := AllData{}
	//point := &dp
	dp.MetricData(MetricsName)
	fmt.Println(dp.DataPoint)

}
