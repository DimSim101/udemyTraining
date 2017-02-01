package main

import (
	"fmt"
)

var counter int
var channel chan int
var done chan bool

func main() {
	counter = 42
	channel = make(chan int)
	done = make(chan bool)
	fmt.Println("Start counter value:", counter)

	go func() {
		channel <- counter // place initial counter value into channel
	}()
	go increment()
	go increment()

	// After both increment() have finished, we should have a single
	// int inside channel which contains the new counter value (incremented).
	go finish()
	//fmt.Println("End counter value:", counter)
}

func increment() {
	for i := 0; i < 10; i++ {
		channelVar := <-channel // get counter value out of channel
		fmt.Println("Incrementing value:", channelVar)
		channelVar++            // increment local value
		channel <- channelVar   // put new value (counter+1) on channel

		// EXPLANATION FOR WHY <-done <-done as the first two lines of
		// finish() results in finish() blocking and main() ending first.

		// WHEN IN THE LAST ROUND OF THE LAST INCREMENT WE CANNOT PUT
		// channelVar ON THE CHANNEL BECAUSE THE RECEIVER FOR IT
		// ISN'T READY (were blocking on the second done in finished()).
		// THUS, IF WE REMOVE THE BLOCK ON THE SECOND DONE, ONCE THE FIRST
		// INCREMENT IS COMPLETE, THEN WE CAN GO AHEAD AND FINISH
		// THE SECOND INCREMENT WHICH IS NO LONGER GETTING STUCK
		// AT THE FINAL STEP. WE CAN MOVE THIS SECOND DONE TO AFTER
		// WE GRAB THIS LAST VALUE OUT OF THE CHANNEL TO ENSURE EVERYTHING
		// IS COMPLETE.
		//
		// HOWEVER, AFTER THE FIRST INCREMENT() FINISHES
		// AND finish() UNBLOCKS THE FIRST <-done, THERE WILL BE ONE
		// INT IN THE CHANNEL. WHAT IS TO STOP COUNTER = <- CHANNEL
		// HERE, UPDATE THE COUNTER AND THEN BLOCK ON <-done again?
		// IF THIS OCCURRED, THEN THE PROGRAM WOULD JUST END BECAUSE
		// FINISH WOULD BE BLOCKING ON <-done AND INCREMENT WILL
		// BE BLOCKING ON channelVar := <-channel AS THERE IS NOTHING
		// INSIDE THE CHANNEL ANYMORE (VALUE WAS TAKEN BY COUNTER SO
		// CHANNEL IS NOW EMPTY).

		// HOW CAN WE TEST FOR THIS?
	}
	done <- true
}

func finish() {
	<-done
	counter = <- channel
	<-done
	close(channel)
	fmt.Println("End counter value:", counter)
	fmt.Println("FINISHED")
}
