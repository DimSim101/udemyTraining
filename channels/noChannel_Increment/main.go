package main

import (
	"fmt"
	"sync"
	"time"
)

var counter int

//var done chan bool
var semaphore sync.WaitGroup
var mutex sync.Mutex

func main() {
	//done = make(chan bool)
	semaphore.Add(2)
	counter = 42
	fmt.Println(counter)
	go increment()
	go increment()
	// go finish()
	semaphore.Wait()
	fmt.Println(counter)
}

func increment() {
	for i := 0; i < 10; i++ {
		time.Sleep(5 * time.Millisecond)
		mutex.Lock()
		time.Sleep(5 * time.Millisecond)
		counter++
		time.Sleep(5 * time.Millisecond)
		mutex.Unlock()
		time.Sleep(5 * time.Millisecond)
	}
	//done <- true
	semaphore.Done()
}

/*
func finish() {
	// This doesn't work, doesn't wait on <-done, just bypasses and ends.
	// Instead, we must use a Semaphore / WaitGroup
	//<-done
	//<-done
	//close(done)
}
*/
