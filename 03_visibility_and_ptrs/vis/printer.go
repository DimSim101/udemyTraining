package vis

import "fmt"

// PrintVar is exported because it starts with a capital letter
func PrintVar() {
	fmt.Println(MyName)
	fmt.Println(yourName)
	fmt.Println(Q)
	fmt.Println(q)
}

// Lower case variables are like "private variables", external classes /
// packages cannot manipulate their values directly (as they don't have
// access to / cannot create a pointer to the variable directly - as can be
// done with Upper case (externally visible) variables.

// GetNamePtr is an externally visible function
// The only case in which a lower case (private) variable can be manipulated
// directly is through returning a pointer to the local variable in some way
// (i.e. a function like GetNamePtr() below)
func GetNamePtr() *string {
	return &yourName
}
