package main

import (
	"fmt"
	//"math/rand"
	"sync"
	"time"
)

var m = sync.RWMutex{} //opposed to the classic Mutex RWMutex also adds the RLock and RUnlock methods for reading
var wg = sync.WaitGroup{}
var dbData = []string{"id1", "id2", "id3", "id4", "id5"}
var results = []string{}

func main() {
	t0 := time.Now()
	for i := range len(dbData) {
		wg.Add(1)
		go dbCall(i)
	}
	wg.Wait()
	fmt.Printf("Total execution time: %v\n", time.Since(t0))
	fmt.Printf("The results are: %v\n", results)
}

func dbCall(i int) {
	var delay float32 = 2000
	time.Sleep(time.Duration(delay) * time.Millisecond)
	save(dbData[i])
	log()
	wg.Done()
}

func save(result string) {
	m.Lock() //writing lock
	results = append(results, result)
	m.Unlock()
}

func log() {
	m.RLock() //reading lock
	fmt.Printf("The current results are: %v\n", results)
	m.RUnlock()
}
