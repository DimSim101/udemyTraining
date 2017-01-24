package main

import "fmt"

func main() {
	// test exercises 1 and 2 with a basic loop
	for i := 0; i < 10; i++ {
		newNum, even := ex1(i)
		fmt.Printf("Example 1: [%v] = (%v, %v) \n", i, newNum, even)

		ex2 := func(num int) (int, bool) {
			return num / 2, num % 2 == 0
		}
		fmt.Println(ex2(i))
	}

	// test ex3
	ex3(1,2,3,4,5,6,7,8,9)
	ex3(9,8,7,6,5,4,3,2,1)
	ex3(1,2,3,4,9,8,7,6,5)
	foo(1,2)
	foo(1,2,3)
	aSlice := []int{1,2,3,4}
	foo(aSlice...)
	foo()
	fmt.Println()
	fmt.Print("Total sum of all even numbers in: ...")
	fmt.Println("EULER RESULT IS:", euler(1, 2, 0))

}

func ex1(number int) (int, bool) {
	/*
	newNum := number / 2
	if number % 2 == 0 {
		return newNum, true
	}
	return newNum, false
	*/
	return number / 2, number % 2 == 0
}


func ex3(numbers ...int) int {
	var largestNum int
	var largestNumIndex int
	for key, value := range numbers {
		if value > largestNum {
			largestNum = value
			largestNumIndex = key
		}
	}

	fmt.Println("Largest number found at index[", largestNumIndex, "] with value: ", largestNum)
	return largestNum
}

func foo(numbers ...int) {
	for _, value := range numbers {
		fmt.Println(value)
	}
}


// function adds up to a limit (in this case 420) all the numbers in the fibonacci
// sequence. The sum of all numbers is simply the second number in the pair.
// If a real sum was required, it is harder to do this recursively as the
// variable would be consistently reset on each new call
// We could alternatively pass around the sum so far which would then solve
// this problem, and I have done so here.
func euler(startOne int, startTwo int, sum int) int{
	var totalEvenSum int

	totalEvenSum += sum

	if startTwo > 420 {
		fmt.Println("=", totalEvenSum)
		return startTwo
	}

	fmt.Print(startOne)
	fmt.Print(" + ")
	if startOne % 2 == 0 {
		totalEvenSum += startOne
	}

	return euler(startTwo, startOne + startTwo, totalEvenSum)
}

