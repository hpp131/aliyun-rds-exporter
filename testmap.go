package main

import "fmt"

func main() {
	fmt.Println("starting...")
	testmap := make(map[int]int)
	for i := 1; i < 1000; i++ {
		testmap[i] = i
	}
	fmt.Println(testmap)
}
