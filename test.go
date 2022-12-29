package main

import (
	"fmt"
)

func main() {

	//var MetricsName = []string{"ConnectionUsage", "CpuUsage", "DiskUsage", "IOPSUsage", "MemoryUsage", "MySQL_ActiveSessions"}
	////var test = []string{}
	////for index, metric := range MetricsName {
	////	fmt.Println(index, metric)
	////	//test = append(test, metric)
	////}
	//fmt.Println(MetricsName[0])

	// 结构体
	type structtest struct {
		name string
		age  int
	}
	//var ming  *structtest
	////ming.age = 12
	////fmt.Println(ming)
	//var hua structtest
	//fmt.Printf("%p\n", hua)
	//hua.age = 12
	//fmt.Printf("%p\n", hua)

	// slice
	//slicetest := make([]string, 5)
	//fmt.Println(slicetest)
	//fmt.Printf("%p\n", slicetest)
	//if slicetest == nil {
	//	fmt.Println("slicetest is nil")
	//}

	//map
	//var maptest  map[int]string
	////maptest = map[int]string{1: "asdfas"}
	//fmt.Println(maptest)
	//fmt.Printf("%p\n", maptest)
	//if maptest == nil {
	//	fmt.Println("maptest is nil")
	//}

	//maptest2 := make(map[int]int)
	//fmt.Println(maptest2)
	//fmt.Printf("maptest2 address is %p\n", maptest2)
	//if maptest2 == nil {
	//	fmt.Println("maptest2 is nil")
	//}

	//new()
	ming := new(structtest)
	fmt.Println(ming)
	testadd
}
