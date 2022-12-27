package main

import "fmt"

func main() {

	var MetricsName = []string{"ConnectionUsage", "CpuUsage", "DiskUsage", "IOPSUsage", "MemoryUsage", "MySQL_ActiveSessions"}
	var test = []string{}
	for _, metric := range MetricsName {
		fmt.Println(metric)
		test = append(test, metric)
	}
	fmt.Println(test)

}
