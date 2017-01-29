package main

import (
	"fmt"
	"encoding/json"
	"log"
)

type Animal struct {
	Category string
	NumLegs int
	Name string
	// as the colour variable is not expored / not externally visible,
	// when we pass data of this type to json.Marshal
	// (which is using an external library file) it is then out of scope
	// and cannot be accessed. Thus, the returned data variable below
	// will not contain json data of the colour variable.
	colour string
	//This Colour variable is externally visible and thus gets included
	Colour string
}

func newAnimal (category string, numLegs int, name string, colour string) Animal {
	return Animal{Category:category,
		NumLegs:numLegs,
		Name:name,
		colour:colour,
		Colour:colour,
		}
}

func main() {
	a := newAnimal("Canine", 4, "Chichi", "Mixed")
	a2 := newAnimal("Canine", 4, "Bowie", "Black")

	l := []Animal{a, a2}

	data, error := json.Marshal(l)

	if error != nil {
		log.Fatal(error)
	}

	fmt.Println(string(data))
}