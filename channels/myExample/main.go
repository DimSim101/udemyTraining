package main

import (
	"fmt"
)

var counter int
var channel chan int
var done chan bool
var tempCounter int

func main() {
	counter = 42
	channel = make(chan int)
	done = make(chan bool)

	go increment()
	go increment()
	go finish()

	for n := range channel {
		fmt.Println(n)
		tempCounter = n // Update tempCounter with new value coming out
		// of channel (newly incremented value)
	}

	counter = tempCounter // update counter value to be the last value out
	// of the channel.
	// We cannot do counter = n inside the for loop above because then we
	// are manipulating counter while potentially calling increment() which
	// uses the counter value to put it inside the channel (channel <- counter).
	// Hence, we can use a temp variable to store the values returned from
	// the channel and update it accordingly. Once the channel is closed and
	// the loop is complete, we can then set counter safely to its newly
	// incremented value.
	fmt.Println(counter) // value remains unchanged as we only passed
	// the value of the counter to the channel
	// I tried using a *int inside the channel but we get a race condition
	// when manipulating the value via *channeVar++. Probably easiest to
	// just use a mutex when manipulating a shared single global int like we did
	// before.

	// NOTE: IM NOT SURE THIS IS THE BEST / A CORRECT SOLUTION AS I FEEL AS
	// THOUGH USING A MUTEX TO LOCK DOWN THE GLOBAL COUNTER IS MUCH EASIER
	// AND 100% SAFE.

	// I'm confused about how this would actually work successfully. We
	// have counter = 42. We call increment. 42 gets put on the channel.
	// Then we grab the value off the channel, increment it and put it
	// on the channel. Lets image we have two increments running concurrently
	// as we do.

	// Now, channel = 42. We increment, and put it on channel. Channel = 43.
	// Now if we call the second increment, since there is something on the
	// channel it must block and wait (it cannot do channel <- counter).
	// Thus, the first increment will start again, this time going into the
	// loop and grabbing the value out of the channel. We then locally increment
	// this value and put it back on the channel. Before we place the new value
	// in the channel, it is possible for the second increment to be called
	// in an attempt to place counter on the channel. Now we can increment
	// and we have channelVar = 43 here too. Now we place 43 on the channel.
	// If we try and go back to the first increment called, we cannot place
	// channelVar on the channel since it now has a value. Thus, the pattern
	// continues and can swap/not swap between increment after taking the value
	// from the channel and storing it in a local var.

	// Lets image for a second that first increment called always wins this
	// "race" and manages to get priority over the channel. This would result
	// in: Counter = 42. Put into channel. Get out of channel, increment,
	// put new value in channel. Channel = 43. If we do this 10 times,
	// channel = 52 and done = true. Now finish() must wait for a second done
	// to the true. Now channel = 52 and increment() is called (for a second
	// time). Increment() will try channel <- counter and block because
	// the channel has a value in it. This will cause the range loop to grab
	// the value out, print it (52), before the counter value (42) can be
	// placed inside the channel. However, this doesn't happen.

	// What seems to happen is: Counter = 42. increment() is called. We place
	// counter in the channel. Now, range has a value inside the channel it
	// can grab. Hence, the very first value in the channel gets printed.
	// Now, if the same increment() is called again, it will start the loop
	// and block  at channelVar := <-channel because nothing is in the channel.
	// Thus, the only option is for increment() to be called again. This time
	// again the counter is placed inside the channel. Its possible that here
	// range loop would print the value and we would deadlock. // This is what
	// happens if we uncomment the sleep line underneath channel <- counter,
	// as we force a context switch. However, this doesn't seem to happen
	// otherwise. Hence, we have now essentially throw away (via printing)
	// the first value placed into the channel (i.e. one of the two counter
	// values passed in when increment() is called). Thus, even though this
	// is super messy and probably will break under different conditions that
	// result in a context switch causing deadlock, this seems to work
	// and will correctly have counter + 20 as the last value in the channel.
}

func increment() {
	channel <- counter // put counter into channel
	//time.Sleep(5*time.Millisecond) // uncomment this for deadlock.
	for i := 0; i < 10; i++ {
		channelVar := <-channel // get counter value out of channel
		channelVar++            // increment local value
		channel <- channelVar   // put new value (counter+1) on channel
	}
	done <- true
}

func finish() {
	<-done
	<-done
	close(channel)
}
