package main

import "fmt"

func main() {
	c := incrementor()
	for n := range puller(c) {
		fmt.Println(n)
	}

	// This is a channel that can have data sent to it and read from it.
	// This type of channel is called bidirectional.
	newChan :=  make(chan int)
	// Once passed in here, the local channel variable within adderOnly()
	// becomes send only (channel can only have stuff sent to it).
	adderOnly(newChan)

	// Since the type didnt change, we can still iterate over it.
	fmt.Println("newChan = ")
	for value := range newChan {
		fmt.Println(value)
	}

	// After iterating over newChan, the values are gone. We repeat the
	// process to show how we can manipulate some channel stuff.
	newChan2 :=  make(chan int)
	adderOnly(newChan2) // 42,43,44,45 again inside channel

	// Here we take a channel which we can only receive from. This changes
	// the local value of channel inside the receiverOnly() function to
	// be receive only. Since it returns a channel of the same type,
	// we can iterate over these values and see they are the same as the
	// original channel.
	newChan3 := receiverOnly(newChan2)

	fmt.Println("newChan3 = ")
	for value := range newChan3 {
		fmt.Println(value)
	}

	// The line below will fail as newChan3 is receive only (we can only
	// receive values from it)
	//newChan3 <- 99
}


// Takes in a channel you can only receive from
// returns a channel you can only receive from
func receiverOnly(channel <-chan int) <-chan int{
	out := make(chan int)

	go func() {
		// Here we can loop over the channel values since we can
		// receive from it. We cannot do channel <- someValue because
		// that would be sending to the channel.
		for val := range channel {
			out <- val
		}
		close(out)
	}()
	return out
}

// takes in a channel that you can only send to.
func adderOnly(channel chan<-int) {
	go func() {
		for i := 42; i < 47; i++ {
			channel<-i
			// Cannot read values below as local channel variable
			// is send only.
			/*
			for value := range channel {
				fmt.Println(value)
			}
			*/
		}
		close(channel)
	}()
}

// returns a channel we can only receive things from.
func incrementor() <-chan int {
	out := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			out <- i
		}
		close(out)
	}()
	return out
}

// takes in a channel we can only receive from and returns a channel we can only
// receive from.
func puller(c <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		var sum int
		for n := range c {
			sum += n
		}
		out <- sum
		close(out)
	}()
	return out
}

/*
The optional <- operator specifies the channel direction, send or receive.
If no direction is given, the channel is bidirectional.
https://golang.org/ref/spec#Channel_types
*/
