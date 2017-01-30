package main

import (
	"fmt"
	"strconv"
)

func main() {

	fmt.Println("Initializing exercise 1...")
	exercise1()

	fmt.Println("Initializing exercise 2...")
	exercise2()

	fmt.Println("Initializing exercise 3...")
	exercise3()

	fmt.Println("Initializing exercise 4...")
	exercise4()

	fmt.Println("Initializing exercise 5...")
	exercise5()

	fmt.Println("Initializing exercise 6...")
	exercise6()

	fmt.Println("Initializing exercise 7...")
	exercise7()

	fmt.Println("All tests passed! You are awesome!")
}

func exercise1() {
	fmt.Println("Hello World")
}

func exercise2() {
	fmt.Println("Hello my name is David")
}

func exercise3() {
	var name string
	fmt.Println("Please enter your name:")
	_, error := fmt.Scan(&name)

	if error == nil {
		fmt.Println("Hello", name)
	}
}

func exercise4() {
	var smallNum int
	var largeNum int
	fmt.Println("Please enter a small number:")
	_, error := fmt.Scan(&smallNum)

	if error == nil {
		fmt.Println("Please enter a large number:")
		_, error := fmt.Scan(&largeNum)
		if error == nil {
			remainder := largeNum % smallNum
			fmt.Println(remainder)
		}
	}
}

func exercise5() {
	for i := 1; i < 100; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}
}

func exercise6() {
	for i := 1; i <= 100; i++ {
		var output string
		if i%3 == 0 {
			output += "Fizz"
		}
		if i%5 == 0 {
			output += "Buzz"
		}

		if output == "" {
			output = strconv.Itoa(i)
		}
		fmt.Println(output)
	}
}

func exercise7() {
	var sum int
	for i := 1; i < 1000; i++ {
		fmt.Println(i)
		if i%3 == 0 || i%5 == 0 {
			sum += i
		}
	}
	fmt.Println(sum)
}
