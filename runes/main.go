package main

import (
	"fmt"
)

func main() {

	fmt.Println()
	// since a rune is an alias for an int32, we must convert characters
	// to strings in order to get their value as a string not an int.
	fmt.Println(`testing runes on rune 'a' without string() conversion:`,
		'a')
	fmt.Println(`testing runes on rune 'a' with conversion:`, string('a'))

	fmt.Println()

	b := "bbb"
	fmt.Println("Multicharacter string created:", b)
	runeB := []rune(b)
	fmt.Println(`Converted into slice of rune's`, runeB)

	s := string(runeB)
	fmt.Println("Converted back into string with contents:", s)
}
