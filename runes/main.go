package main

import (
	"fmt"
)

func main() {
	// since a rune is an alias for an int32, we must convert characters
	// to strings in order to get their value as a string not an int.
	fmt.Printf("testing runes" + string('a'))
}
