package main

import (
	"testing"
	"fmt"
	"strconv"
	"sort"
)

var expectedOutput = make(map[int]string)
var incorrectOutput = make(map[int]string)

func TestFizzBuzz(t *testing.T) {

	fmt.Println("SETTING UP TESTS WITH THE FOLLOWING KEY - VALUE PAIRS")
	initTest()

	fmt.Println("RUNNING TESTS!!!")
	for i := 1; i <= 50; i++ {
		var invalidMatch bool
		actualOutput := calcFizzBuzz(i)
		if actualOutput != expectedOutput[i] {
			t.Errorf("Expected: %v -- Got: %v -- NO MATCH!\n", expectedOutput[i], actualOutput)
			invalidMatch = true
		}

		if actualOutput == incorrectOutput[i] {
			t.Errorf("Incorrect: %v -- Got: %v -- INCORRECT TEST MATCH!\n", incorrectOutput[i], actualOutput)
			invalidMatch = true
		}

		if !invalidMatch {
			fmt.Printf("TEST PASSED!!! Expected: %v -- Got %v -- SUCCESSFULL MATCH!\n", expectedOutput[i], actualOutput)
		}
	}
}


func initTest() {

	for i := 1; i <= 50; i++ {
		if i % 3 == 0 && i % 5 == 0 {
			expectedOutput[i] = "FizzBuzz"
			incorrectOutput[i] = "fizzbuzz"
		} else if i % 3 == 0 {
			expectedOutput[i] = "Fizz"
			incorrectOutput[i] = "Buzz"
		} else if i % 5 == 0 {
			expectedOutput[i] = "Buzz"
			incorrectOutput[i] = "Fizz"
		} else {
			expectedOutput[i] = strconv.Itoa(i)
			incorrectOutput[i] = strconv.Itoa(i+1)
		}
	}

	//PRINT TEST TO CHECK VALUES USED IN EXPECTED RESULTS

	var keys []int
	for k := range expectedOutput {
		keys = append(keys, k)
	}

	sort.Ints(keys)

	fmt.Println("CORRECT EXPECTED RESULTS: -----------")
	for _, k := range keys {
		fmt.Println("Key:", k, "Value:", expectedOutput[k])
	}

	fmt.Println("INCORRECT RESULTS FOR TESTING: -----------")


	var incorrectKeys []int
	for k := range incorrectOutput {
		incorrectKeys = append(incorrectKeys, k)
	}

	sort.Ints(incorrectKeys)

	for _, k := range incorrectKeys {
		fmt.Println("Key:", k, "Value:", incorrectOutput[k])
	}
}

func BenchmarkFizzBuzz(*testing.B) {

}
