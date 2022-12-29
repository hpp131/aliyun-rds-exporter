package main

import "fmt"

func main() {

	var MetricsName = []string{"ConnectionUsage", "CpuUsage", "DiskUsage", "IOPSUsage", "MemoryUsage", "MySQL_ActiveSessions"}
	//var test = []string{}
	//for index, metric := range MetricsName {
	//	fmt.Println(index, metric)
	//	//test = append(test, metric)
	//}
	fmt.Println(MetricsName[0])
}
