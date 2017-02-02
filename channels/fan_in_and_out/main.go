package main

import (
	"fmt"
)

func main() {
	c := fanIn(boring("David"), boring("Mandy"))
	for i := 0; i < 10; i++ {
		fmt.Println(<-c)
	}
	fmt.Println("Your boring im peacing it...")
}

func boring(name string) <-chan string {
	out := make(chan string)
	go func() {
		for i := 0; i < 5; i++ {
			out <- fmt.Sprintf("%s - %d", name, i)
		}
	}()
	return out
}

func fanIn(input1, input2 <-chan string) <-chan string {
	out := make(chan string)
	go func() {
		for {
			out <- <-input1
		}
	}()

	go func() {
		for {
			out <- <-input2
		}
	}()
	return out
}
