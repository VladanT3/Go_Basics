package main

import (
	"fmt"
	"sync"
	"time"
)

var wg = sync.WaitGroup{}

func main() {
	t0 := time.Now()
	for range 1000 {
		wg.Add(1)
		go count()
	}
	wg.Wait()
	fmt.Printf("Total execution time: %v\n", time.Since(t0))
}

func count() {
	var res int
	for i := 0; i < 100000000; i++ {
		res += 1
	}
	wg.Done()
}
