package main

import (
	"fmt"
	"encoding/json"
	"log"
)

type animal struct {
	// Here I have added some tags to show their usage in an example
	// and also make use of them by cleaning up some variable names so
	// the printed example json reads nicely.
	Category string `json:"Type"`
	NumLegs int `json:"Number of legs"`
	Name string
	// as the colour variable is not expored / not externally visible,
	// when we pass data of this type to json.Marshal
	// (which is using an external library file) it is then out of scope
	// and cannot be accessed. Thus, the returned data variable below
	// will not contain json data of the colour variable.
	colour string
	//This Colour variable is externally visible and thus gets included
	Colour string
	// The `json:"-"` syntax below says not to include this variable in json
	VisibleButExcludedColour string `json:"-"`
}

func newAnimal (category string, numLegs int, name string, colour string) animal {
	return animal{Category:category,
		NumLegs:numLegs,
		Name:name,
		colour:colour,
		Colour:colour,
		VisibleButExcludedColour:colour,
		}
}

func main() {
	a := newAnimal("Canine", 4, "Chichi", "Mixed")
	a2 := newAnimal("Canine", 4, "Bowie", "Black")

	l := []animal{a, a2}

	data, error := json.Marshal(l)

	if error != nil {
		log.Fatal(error)
	}

	fmt.Println(string(data))
}