package main

import (
	"fmt"
	"sync"
	"time"
	"math/rand"
)

// This is a function that initializes a count to 0 and returns a function.
// This way we can init the variable and manipulate it with the returned function.
// This means we only call wrapper() once, and therefor x is initialized to its
// default 0 once (when called) and then incremented from there using the returned
// function. See below for example.
func wrapper() func() int {
	var x int
	return func() int {
		x++
		return x
	}
}

func concurrencyWrapper() func() {
	var counter int
	var mutex sync.Mutex
	return func () {
		for i := 0; i < 10; i++ {
			time.Sleep(time.Duration(rand.Intn(20))*time.Millisecond)
			mutex.Lock()
			counter++
			time.Sleep(time.Duration(rand.Intn(20))*time.Millisecond)
			fmt.Println("Counter value is:", counter)
			mutex.Unlock()
		}
		semaphore.Done()
	}
}

var semaphore sync.WaitGroup
func main() {
	// Here we get the function result that is returned from wrapper(),
	// and calling wrapper() also initializes the int which increment() accesses.
	// Thus, calling increment() only calls the inner function and doesn't
	// reset the value of x.
	increment := wrapper()
	fmt.Println(increment())
	fmt.Println(increment())

	// Here we can use this function wrapper() concept to initalize a mutex
	// and a counter together which can then be manipulated using the
	// function returned from concurrencyWrapper().
	// We still use a global semaphore (WaitGroup) object so that main and
	// concurrencyWrapper() both can access it within their scope.
	// This is a nice example of how we can use functions, scope and closure
	// to restrict variables to certain functions / combine variables together
	// in a way that provides access and encapsulation with their scope.
	concurrentInc := concurrencyWrapper()
	semaphore.Add(2)
	go concurrentInc()
	go concurrentInc()
	semaphore.Wait()
}

/*
closure helps us limit the scope of variables used by multiple functions
without closure, for two or more funcs to have access to the same variable,
that variable would need to be package scope
*/
