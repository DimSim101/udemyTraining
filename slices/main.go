package main

import "fmt"

func main() {

	var results []int
	fmt.Println(results)

	mySlice := []string{"a", "b", "c", "g", "m", "z"}
	fmt.Println(mySlice)
	fmt.Println(mySlice[2:4])  // slicing a slice
	fmt.Println(mySlice[2])    // index access; accessing by index
	fmt.Println("myString"[2]) // index access; accessing by index
	fmt.Println(string("myString"[2])) // index access; get original value out.

	// make: allocates a slice of a given length as well as an underlying
	// array of length either: == to len(slice) or == cap (3rd arg to make)
	// make will return a slice which can be allocated up to its index(len()-1)
	// as well as an array of given capacity / same length as the slice.
	// The slice can be filled using the normal syntax slice[blah] = blah
	// However, if the slice index used does not exist within the current
	// slice (index used > len(slice) - 1) then its out of bounds.
	// Thus, in order to safely add to a slice, we can use append (or other
	// built in functions depending on what we want to do) and reassign
	// the slice to the result.

	// This way the item will be added to the slice, and the slices length will
	// be updated accordingly. Additionally, if the added item results in the
	// slices length exceeding its capacity, a new array is allocated with
	// double the size and the slice is stored there instead (allowing
	// for dynamic sizes for slices during runtime). However, after a certain
	// size limit, the new array is no longer doubled in size but instead
	// increased by some small amount (to stop taking up increasingly
	// larger chunks when they are probably not needed).
	customerNumber := make([]int, 3)

	fmt.Println(customerNumber)
	fmt.Printf("LEN IS: %v and CAP IS: %v\n", len(customerNumber), cap(customerNumber))
	// 3 is length & capacity
	// // length - number of elements referred to by the slice
	// // capacity - number of elements in the underlying array
	customerNumber[0] = 7
	customerNumber[1] = 10
	customerNumber[2] = 15

	fmt.Println(customerNumber)
	fmt.Printf("LEN IS: %v and CAP IS: %v\n", len(customerNumber), cap(customerNumber))

	// line below fails
	// customerNumber[3] = 420

	customerNumber = append(customerNumber, 420)
	fmt.Println(customerNumber)
	fmt.Printf("LEN IS: %v and CAP IS: %v\n", len(customerNumber), cap(customerNumber))

	greeting := make([]string, 3, 5)
	// 3 is length - number of elements referred to by the slice
	// 5 is capacity - number of elements in the underlying array
	// you could also do it like this

	fmt.Printf("Greeting = %v\n", greeting)
	fmt.Printf("LEN IS: %v and CAP IS: %v\n", len(greeting), cap(greeting))

	greeting[0] = "Good morning!"
	greeting[1] = "Bonjour!"
	greeting[2] = "dias!"
	// line below fails
	//greeting[3] = "test"
	// append works. Will change the slice and its length but keep the
	// underlying array (cap) if it can fit
	// Otherwise, appears to just allocate a new array of size (cap*2)
	// and presumably free the old one after storing the data to it.

	fmt.Printf("Greeting = %v\n", greeting)
	fmt.Printf("LEN IS: %v and CAP IS: %v\n", len(greeting), cap(greeting))

	greeting = append(greeting, "test")
	fmt.Println("Appended string: test")
	fmt.Printf("Greeting = %v\n", greeting)
	fmt.Printf("LEN IS: %v and CAP IS: %v\n", len(greeting), cap(greeting))

	greeting = append(greeting, "test2")
	fmt.Println("Appended string: test2")
	fmt.Printf("Greeting = %v\n", greeting)
	fmt.Printf("LEN IS: %v and CAP IS: %v\n", len(greeting), cap(greeting))

	greeting = append(greeting, "test3")
	fmt.Println("Appended string: test3")
	fmt.Printf("Greeting = %v\n", greeting)
	fmt.Printf("LEN IS: %v and CAP IS: %v\n", len(greeting), cap(greeting))

	// This works the same way as make (must call make somewhere when we
	// initialize a slice) and makes a slice of len == 0 and cap == 0
	// because the pointer to the underlying array is nil (unitialized default value).
	var testSlice []string
	fmt.Printf("testSlice is: %v - LEN IS: %v and CAP IS: %v\n", testSlice, len(testSlice), cap(testSlice))

	// After adding something to it, the same thing applies as above
	// Add the item if possible and increase length by 1. If array is full
	// double array size, allocate a new array, copy all items to new array,
	// free old array and then allocate the item to the new array.
	testSlice = append(testSlice, "aaa")
	fmt.Println("Appending aaa to testSlice")
	fmt.Printf("testSlice is: %v - LEN IS: %v and CAP IS: %v\n", testSlice, len(testSlice), cap(testSlice))

	testSlice = append(testSlice, "bbb")
	fmt.Println("Appending bbb to testSlice")
	fmt.Printf("testSlice is: %v - LEN IS: %v and CAP IS: %v\n", testSlice, len(testSlice), cap(testSlice))

	testSlice = append(testSlice, "ccc")
	fmt.Println("Appending ccc to testSlice")
	fmt.Printf("testSlice is: %v - LEN IS: %v and CAP IS: %v\n", testSlice, len(testSlice), cap(testSlice))

	// Note: Regardless of the above, it is important to realize the purpose
	// of functions like make (and new etc.)
	// If we are initializing some slice, say a slice of strings, which we know
	// needs 50 values at initialization, but will later need another 50.
	// Instead of just creating an empty slice of length 0, and then
	// adding items and having the array created and doubled in size many
	// times, we can create a slice with a length 50 and a capacity 100.
	// We can then store all our items easily, and we arent wasting
	// efficiency / time creating and freeing pointless arrays as we double
	// in size to allocate all the items repeatedly.

	// See the example here:
	madeSlice := make([]int, 50, 100)
	fmt.Printf("madeSlice is: %v - LEN IS: %v and CAP IS: %v\n", madeSlice, len(madeSlice), cap(madeSlice))

	for i := 0; i < len(madeSlice); i++ {
		madeSlice[i] = i
	}

	// Notice both len and cap remain unchanged, as we have neither exceeded
	// initial 50 values in the slice, nor needed to allocate a new array
	fmt.Printf("madeSlice is: %v - LEN IS: %v and CAP IS: %v\n", madeSlice, len(madeSlice), cap(madeSlice))

	// Another example of where this would be useful is if ints represented
	// some state, of which we needed initialized to be 0.
	// For example, when we created our game board for the AI simple maze
	// text navigation game, we needed to initialize a list of values to be
	// 0, which would then change based on what the AI had seen each turn
	// (The blank 0's are the representation of the game board in the data
	// when the game first begins).
	// Thus, dynamically allocating a slice using var slice []int to init
	// the slice and then append the 0's into it would again result in
	// having the double the underlying array repeatedly to fit all the items
	// Instead, we can create the size needed and initialize the values using
	// make in one go, making the overall procedure much more efficient.


	intSlice := make([]int, 0, 3)

	fmt.Println()
	fmt.Println("-----------------")
	fmt.Println("Len:", len(intSlice), "Capacity:", cap(intSlice), "Value: ", intSlice)
	fmt.Println("-----------------")
	fmt.Println()

	for i := 0; i < 7; i++ {
		// Here we print the address of the pointer to the slice
		// before and after appending so we can see that the address
		// changes when length exceeds capacity and a new underlying
		// array is created to store the slice items.
		fmt.Printf("INTSLICE ADDRESS = %p \n", intSlice)
		fmt.Printf("This is the address of the slice itself (designated by the first elem)\n\n")
		fmt.Printf("INTSLICE POINTER ADDRESS = %p \n", &intSlice)
		fmt.Printf("This is the address of the pointer to the underlying array of items" +
			"that make up a slice. This never changes (even after" +
			"array is reallocated etc (which makes sense)).\n\n")
		intSlice = append(intSlice, i)
		fmt.Printf("Address of first elem (post append) is %p \n", &intSlice[0])
		fmt.Printf("Address of intSlice after append is %p \n", intSlice)
		fmt.Println("Len:", len(intSlice), "Capacity:", cap(intSlice), "Value: ", intSlice[i])
		fmt.Println()
	}

	fmt.Println()
	fmt.Println("------------------")
	fmt.Println("Deleting from slice:")
	fmt.Println()

	oneToFive := []int{1,2,3,4,5}
	fmt.Println(oneToFive)

	oneToFive = deleteMiddleElem(oneToFive)
	//for i := 0; i < len(oneToFive); i++ {
	//	// delete the middle element
	//	if (i == len(oneToFive)/2) {
	//		fmt.Println("deleting middle element")
	//		oneToFive = append(oneToFive[:i], oneToFive[i+1:]...)
	//	}
	//}
	fmt.Println(oneToFive)
	oneToFive = deleteMiddleElem(oneToFive)
	fmt.Println(oneToFive)
	oneToFive = deleteMiddleElem(oneToFive)
	fmt.Println(oneToFive)
	oneToFive = deleteMiddleElem(oneToFive)
	fmt.Println(oneToFive)
	oneToFive = deleteMiddleElem(oneToFive)
	fmt.Println(oneToFive)
	oneToFive = deleteMiddleElem(oneToFive)
	fmt.Println(oneToFive)

	fmt.Println()

	oneToFive = []int{1,2,3,4,5}
	fmt.Println(oneToFive)
	removeCenter(&oneToFive)
	fmt.Println(oneToFive)
	removeCenter(&oneToFive)
	fmt.Println(oneToFive)
	removeCenter(&oneToFive)
	fmt.Println(oneToFive)
	removeCenter(&oneToFive)
	fmt.Println(oneToFive)
	removeCenter(&oneToFive)
	fmt.Println(oneToFive)
	removeCenter(&oneToFive)
	fmt.Println(oneToFive)
}

// manipulate the given slice using its pointer (more efficient imo)
func removeCenter(slicePtr *[]int) {
	// Since we can call removeCenter on any pointer to a slice of int's
	// we must first check it has some elements before we try remove any
	// otherwise appending (*slicePtr)[halfWay+1] results in index [1]
	// which is out of bounds on an empty slice.
	// While this is an O(1) algorithm for deleting the middle element
	// we need to be more cautious due to the method itself.
	fmt.Println("SLICEPTR BEFORE DELETE = ", *slicePtr)
	if len(*slicePtr) > 0 {
		fmt.Println("LEN =", len(*slicePtr))
		halfWay := len(*slicePtr)/2
		fmt.Println("HALWAY =", halfWay)

		// Why does the line below fail but the next one does not?
		// We expect this line to fail, as lets say we are deleting
		// the last element (index 0), the syntax for this is
		// append((*slicePtr)[:0], (*slicePtr)[0+1:]...)
		// which then becomes append((*slicePtr)[:0], (*slicePtr)[1:]...)
		// where (*slicePtr)[1] is clearly out of range.
		// Thus, if we look at this index directly (using the line below)
		// we expect to get an error about out of bounds (as we have
		// extended beyond the length of the slice).
		// Yet for some reason, the line below throwsm an error,
		// but when we use slicing of slices it does not. Why?
		//fmt.Println("SLICEPTR HALFWAY =", (*slicePtr)[halfWay+1])

		// This one doesn't throw an error, must be something to do
		// with the slicing, but what? - The only difference at all is
		// the existence of the semicolon.
		// i.e. from (*slicePtr)[halfWay+1]) to (*slicePtr)[halfWay+1:])
		// will happily remove the error
		fmt.Println("SLICEPTR HALFWAY =", (*slicePtr)[halfWay+1:])
		*slicePtr = append((*slicePtr)[:halfWay], (*slicePtr)[halfWay+1:]...)
	}
}

// returning a slice of ints
func deleteMiddleElem(slicePtr []int)  []int{
	var newSlicePtr []int

	// O(1) middle elem delete :^)
	if len(slicePtr) > 0 {
		halfWay := len(slicePtr)/2
		newSlicePtr = append(slicePtr[:halfWay], slicePtr[halfWay+1:]...)
	}

	// Simple loop algorithm (O(n)) to delete middle elem
	// Unlike the above algorithms used in deleteMiddleElem() and
	// removeCenter(), because this uses a for loop to iterate over the
	// slice of ints, when the slice is empty, len(slicePtr) == 0,
	// which results in i := 0; i < 0; i++ on the for loop, so we pass
	// over it completely and return the empty slicePtr that was created
	// locally. This isn't great either and I prefer the check on length
	// method used instead.

	//for i := 0; i < len(slicePtr); i++ {
	//	// delete the middle element
	//	if (i == len(slicePtr)/2) {
	//		fmt.Println("deleting middle element")
	//		slicePtr = append(slicePtr[:i], slicePtr[i+1:]...)
	//		// could break here to make algo faster, but still crap
	//	}
	//}
	//fmt.Println(len(slicePtr))
	return newSlicePtr
}
