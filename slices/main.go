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
	// for dynamic sizes for slices during runtime).
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
}

