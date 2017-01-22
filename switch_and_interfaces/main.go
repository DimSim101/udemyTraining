package main

import (
	"fmt"
	"reflect"
)

//  switch on types
//  -- normally we switch on value of variable
//  -- go allows you to switch on type of variable

type contact struct {
	greeting string
	name     string
}

// SwitchOnType works with interfaces
// we'll learn more about interfaces later
func SwitchOnType(x interface{}) {

	val := reflect.ValueOf(x)
	valType := reflect.TypeOf(x).Kind()
	fmt.Println("Passed in value is", val, "with type", valType)

	// could use the below instead of val if we only needed the value when
	// x is of type int32
	runeVal, ok := x.(int32)

	switch x.(type) { // this is an assert; asserting, "x is of this type"
	case int:
		fmt.Println("popped case int")
	case string:
		fmt.Println("popped case string")
	case contact:
		fmt.Println("popped case contact")
	//case int32:
	//	fmt.Println("int32")
	default:
		fmt.Println("popped case unknown")
		if valType == reflect.Int32 {
			fmt.Println("rune with value", val)
		}

		if ok {
			fmt.Println("value is an int32 with value", runeVal)
		}
	}
	fmt.Println()
}

func main() {
	SwitchOnType(7)
	SwitchOnType("McLeod")
	var t = contact{"Good to see you,", "Tim"}
	SwitchOnType(t)
	SwitchOnType(t.greeting)
	SwitchOnType(t.name)
	SwitchOnType('a')
}
