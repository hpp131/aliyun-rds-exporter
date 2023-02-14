package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	test_slice := []int{12, 11, 22, 33}
	for _, value := range test_slice {
		wg.Add(1)
		value := value
		go func() {
			defer wg.Done()
			//time.Sleep(time.Second)
			fmt.Println(value)
		}()
	}
	wg.Wait()
}
