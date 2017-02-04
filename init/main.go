package main

import "fmt"

var number int

// We can have multiple init functions within a single file.
// Each init is called sequentially (as expected). Hence, printed value
// for number will be 420. If you comment out 3rd init(), printed value for
// number will be 7, and if you comment out 2nd and 3rd init(), printed value
// for number will be 42.

func init() {
	number = 42
}

func init() {
	number = 7
}

func init() {
	number = 420
}

func main() {
	fmt.Println(number)
}
