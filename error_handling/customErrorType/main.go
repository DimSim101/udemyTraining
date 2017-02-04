package main

import "fmt"

type factorialError struct {
	invalidValue int
	caller string
	err error
}

func (f factorialError) Error() string {
	return fmt.Sprintf("FactorialError occured. Reason: %v - Caller: %v - Value: %v",
	f.err, f.caller, f.invalidValue)

}

func main() {
	for i := -5; i < 5; i++ {
		fmt.Println("Calling factorial() on:", i)
		val, err := factorial(i, "FOO")
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("Factorial result is:", val)
		}
	}

	number := 42
	fmt.Println("Calling factorial() on:", number)
	val, err := factorial(number, "BLAH")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Factorial result is:", val)
	}


}

func factorial(num int, callerID string) (int, error) {
	var errReason error
	if num < 0 {
		errReason = fmt.Errorf("Cannot call factorial() on negative number: %v", num)
		return 0, factorialError{num, callerID, errReason}
	} else if num > 40 {
		errReason = fmt.Errorf("Cannot call factorial() on such a large value (%v) with type int", num)
		return 0, factorialError{num, callerID, errReason}
	}

	result := 1
	for i := num; i > 0; i-- {
		result *= i
	}
	return result, nil
}