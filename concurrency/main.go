package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"sync/atomic"
	"time"
)

var semaphore sync.WaitGroup
var mutex sync.Mutex
var counter int
var atomicCounter int64

// Not sure if this is still needed? May as well put it in to be 100% were using
// all cpu's available.
func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func main() {
	start := time.Now()

	// We can run go run -race main.go to check for race conditions
	// Each of the chunks below (add/dec and add/add) result in race conditions
	// if we remove the mutex.
	fmt.Println("Example adding 10 and removing 10 from global counter...")
	semaphore.Add(2)
	go increment()
	go decrement()
	semaphore.Wait()
	fmt.Println("Final counter value:", counter)
	fmt.Println()

	fmt.Println("Example adding 10 twice to global counter...")
	counter = 0
	semaphore.Add(2)
	go increment()
	go increment()
	semaphore.Wait()
	fmt.Println("Final counter value is:", counter)
	fmt.Println()

	fmt.Println("Example using counter++ and calling increment twice...")
	counter = 0
	semaphore.Add(2)
	go newInc()
	go newInc()
	semaphore.Wait()
	fmt.Println("Final counter value is:", counter)
	fmt.Println()

	semaphore.Add(2)
	go atomicInc()
	go atomicInc()
	semaphore.Wait()
	fmt.Println("Final atomicCounter value is:", atomicCounter)
	fmt.Println()

	elapsed := time.Since(start)
	fmt.Println("Program took", elapsed)
}

func increment() {
	for i := 0; i < 10; i++ {
		time.Sleep(time.Duration(rand.Intn(20)) * time.Millisecond)
		mutex.Lock()
		// The lines below would be changed to: counter++ instead of
		// separating the increment of the global counter. (see newInc())
		// However, it is left as separate to show concurrency and issues
		// with race conditions. Remove the mutex and its lock/unlock's
		// and this program returns to having a race condition.
		x := counter
		fmt.Println("Start inc: X = ", x, "and counter =", counter)
		x++
		fmt.Println("After inc: X = ", x, "and counter =", counter)
		counter = x
		fmt.Println("End inc: X = ", x, "and counter =", counter)
		mutex.Unlock()
	}
	semaphore.Done()
}

func decrement() {
	for i := 0; i < 10; i++ {
		time.Sleep(time.Duration(rand.Intn(20)) * time.Millisecond)
		mutex.Lock()
		x := counter
		fmt.Println("Start dec: X = ", x, "and counter =", counter)
		x--
		fmt.Println("After dec: X = ", x, "and counter =", counter)
		counter = x
		fmt.Println("End dec: X = ", x, "and counter =", counter)
		mutex.Unlock()
	}
	semaphore.Done()
}

// Running go run -race main.go without the mutex will show this function has a
// race condition as the counter variable is still being accessed at the same
// time (potentially). However, this will still increment correctly as the
// counter++ is a single instruction and cannot context switch during
// its execution. Thus, counter will either not have incremented and it will
// context switch, or counter will have incremented and then a context switch
// can occur or not. Either way, regardless of the race condition the value of
// counter is correct. This is not best practice however. Hence, we have used
// a mutex here to avoid this race condition. The increment() function above
// shows how we can have a race condition within the function itself (by not
// locking critical regions and allowing context switching between instructions)
// as well as between different function calls (always due to context switching).
// In this example, since there is only one instruction the same issue of critical
// region doesn't apply. However, a mutex (lock) must still be used to avoid
// thread switching (context switching) and causing a race condition with counter.
func newInc() {
	for i := 0; i < 10; i++ {
		time.Sleep(time.Duration(rand.Intn(20)) * time.Millisecond)
		mutex.Lock()
		counter++
		mutex.Unlock()
	}
	semaphore.Done()
}

// Since this is a true atomic instruction (cannot be separated into smaller pieces)
// this increment does not require a mutex to avoid the race condition.
// Instead, we use atomic.AddInt64 which takes a pointer to the value we want
// to add to, and a value to increase by. Here we provide 1 as we want to do ++.
func atomicInc() {
	for i := 0; i < 20; i++ {
		time.Sleep(time.Duration(rand.Intn(20)) * time.Millisecond)
		atomic.AddInt64(&atomicCounter, 1)
		// We sleep here to allow for context switch to show only the
		// atomic instruction is thread safe, the prints etc. are not.
		// If we do: go run -race main.go with these lines uncommented,
		// we will get a race condition on atomicCounter due to the print.
		//time.Sleep(time.Duration(rand.Intn(20))*time.Millisecond)
		//fmt.Println(i, "atomicCounter:", atomicCounter)
	}
	semaphore.Done()
}
