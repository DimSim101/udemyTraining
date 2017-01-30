package main

import (
	"fmt"
	"log"
	"strconv"
)

func main() {

	fmt.Println("Starting CONVERSION example...")
	// CONVERSION is used to interchange between types.

	var x int
	var y float64
	x = 12
	y = 12.9230123
	// CONVERSION from float64 to int (drops everything from decimal onwards)
	fmt.Println(int(y) + x)

	var r rune
	var r2 int32 // rune is alias for int32 so this is the same
	r = 'a'
	r2 = 'b'
	// CONVERSION from rune to string.
	fmt.Println(string(r) + string(r2))

	var s []byte
	// each of these is a rune / int32 which gets converted to type byte.
	s = append(s, 'h')
	s = append(s, 'e')
	s = append(s, 'l')
	s = append(s, 'l')
	s = append(s, 'o')
	fmt.Println(string(s))
	// CONVERSION from []byte (slice of byte) to string.
	fmt.Println(string(s))

	// CONVERSION from string to []byte (slice of byte)
	fmt.Println([]byte("Testing"))

	var myInt int
	myInt = 97
	// CONVERSION from int to string (in UTF-8 FORM) i.e. ascii chars.
	fmt.Println(string(myInt))
	// CONVERSION from int to string. Just returns a string containing the int.
	fmt.Printf("value: %s - %d \n", strconv.Itoa(myInt), myInt)
	// CONVERSION from string to int. Just removes quotations to get int.
	value, error := strconv.Atoi("42")
	if error != nil {
		log.Fatal(error)
	}
	fmt.Println(value)

	var myString string
	myString = "a"
	// CONVERSION from string to []byte. Since a string can consist of
	// more than one character (each of which is UTF-8 encoded taking up
	// to 4 bytes), it cannot be represented by an int. Instead, we must
	// convert the string to a slice of bytes each of which contains the
	// underlying byte/int values for the characters inside the string.
	fmt.Println([]byte(myString))
	fmt.Println()

	fmt.Println("Starting ASSERTION example...")
	// ASSERTION is used by interfaces only to assert (essentially convert)
	// the interface to be able to access the concrete type and its value.

	// Here we create a variable (name) of type interface{}
	// This variable has a CONCRETE TYPE of string, and thus can be
	// ASSERTED to be a string via name.(string) syntax which can be used
	// to access the CONCRETE type's value.
	var name interface{} = "Sydney"
	fmt.Printf("Type of name: %T - Type of &name %T\n", name, &name)
	// ASSERTION from interface{} to string.
	str, ok := name.(string)
	if ok {
		fmt.Printf("%T-%q\n", str, str)
	} else {
		fmt.Println("Not a string")
	}

	name = 42
	fmt.Printf("Type of name: %T - Type of &name %T\n", name, &name)
	// ASSERTION from interface{} to string using an int. Will fail.
	str, ok = name.(string)
	if ok {
		fmt.Printf("%T-%q\n", str, str)
	} else {
		fmt.Println("Value is not a string")
	}

	var val interface{} = 7
	// print will show val as type int.
	fmt.Printf("val is of type %T\n", val)
	// yet we cannot do as below, fails on type conversion between int and interface{}
	// fmt.Println(val + 6)
	// Instead we must ASSERT the interface is of type int.
	fmt.Println(val.(int) + 6)

	// If we try CONVERSION, it will fail - cannot convert interface{} to type int
	// - need assertion
	// fmt.Println(int(val) + 6)
}
