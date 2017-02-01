package main

import (
	"crypto/rand"
	"fmt"
)

var counter int

func main() {
	counter = 42
	fmt.Println("Counter =", counter)

	incValues := randomSliceByte(2) // generate 2 random ints (a slice of
	// random bytes of length 2).
	fmt.Println(incValues)

	fmt.Println("Incrementing counter by", incValues[0])
	c1 := incrementByNumber(int(incValues[0]))
	fmt.Println("Incrementing counter by", incValues[1])
	c2 := incrementByNumber(int(incValues[1]))
	c3 := newCount(c1)
	c4 := newCount(c2)
	newCountAddition := <-c3 + <-c4
	counter += newCountAddition
	fmt.Println("Final Counter:", counter)
}

func incrementByNumber(number int) chan int {
	out := make(chan int)
	// Create a channel of ints containing a specified number of 1's
	// so that we can increment our count by a value.

	go func() {
		for i := 0; i < number; i++ {
			out <- 1
		}
		close(out)
	}()
	return out
}

func newCount(c chan int) chan int {
	out := make(chan int)
	go func() {
		var newCount int
		for n := range c {
			newCount += n
		}
		out <- newCount
		close(out)
	}()
	return out
}

func randomSliceByte(size int) []byte {
	if size <= 0 {
		return []byte{}
	}

	b := make([]byte, size)
	_, error := rand.Read(b)
	if error != nil {
		fmt.Println(error)
	}
	return b
}
