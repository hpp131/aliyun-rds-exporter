package main

import (
	"fmt"
)

var testslice []int

func main() {
	for i := 1; i < 10; i++ {
		sliceWrite(testslice)
	}
	fmt.Println(testslice)
}

func sliceWrite(s []int) {
	s = append(s, 1)
}
