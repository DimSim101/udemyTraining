package main

import (
	"fmt"
)

var counter int
var channel chan int
var done chan bool

func main() {
	counter = 42
	channel = make(chan int, 1)
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
}

func increment() {
	for i := 0; i < 10; i++ {
		channelVar := <-channel // get counter value out of channel
		fmt.Println("Incrementing value:", channelVar)
		channelVar++          // increment local value
		channel <- channelVar // put new value (counter+1) on channel

		// Explanation for why <-done <-done as the first two lines of
		// finish() results in finish() blocking and main() ending first.

		// When in the last round of the last increment() call we cannot put
		// channelVar on the channel because the receiver for it
		// isn't ready (were blocking on the second <-done in finished()
		// and so the receiver for <-channel in finished isn't ready yet).

		// We can change the ordering to grab the last value out of the int channel
		// before we block on the second done, bypassing this problem. However,
		// this doesn't ensure that after the first increment finishes and
		// the first <-done passes in finish(), that the program won't just
		// continue inside finish(), grab the last value from the
		// channel of ints, load it into counter and then the next
		// increment() round will be stuck forever trying to receive
		// from an empty channel, resulting in main() just ending.

		// Instead we can use a buffered channel which won't block
		// on putting values inside the channel until its full, and wont
		// block on pulling values out unless its empty. In this case,
		// since we know at the handover point between increment() calls
		// the sender which then blocks as no receiver is ready, we can
		// avoid this block by instead blocking when full (in this case
		// we dont need more than 1 thing inside the channel
		// (similar to unbuffered except we need it to not block)
		// so we can use buffered channel of max size 1).
		// Now, when we get to the end of the last increment and go to
		// put the final value on the buffer, it wont block due to no
		// receiver, as the channel has room for it (will be empty).
		// Thus, this final line executes, true is added to the bool
		// channel, finished() then unblocks, signifying we have finished
		// incrementing our variable. Now we can grab the last value in
		// the channel (which is ready to be received), update the counter
		// and close the channel before ending!
		//BOOOOOOM!
		// I AM DOING GREAT AND I AM AWESOME!
	}
	done <- true
}

func finish() {
	<-done
	<-done
	counter = <-channel
	close(channel)
	fmt.Println("End counter value:", counter)
	fmt.Println("FINISHED")
}
