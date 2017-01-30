package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
)

type animal struct {
	// Here I have added some tags to show their usage in an example
	// and also make use of them by cleaning up some variable names so
	// the printed example json reads nicely.
	// These tags apply to both changing the value when marshaling (to get
	// some data in json form) as well as when unmarshalling json data (as
	// now the data is misslabelled and needs to be identified).
	// For example, for Category below, the json data is "Type":"BLAH"
	// Hence, the program needs some way to trace from Type -> Category.
	Category string `json:"Type"`
	NumLegs  int    `json:"Number of legs"`
	Name     string
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

func newAnimal(category string, numLegs int, name string, colour string) animal {
	return animal{Category: category,
		NumLegs: numLegs,
		Name:    name,
		colour:  colour,
		Colour:  colour,
		VisibleButExcludedColour: colour,
	}
}

func main() {

	fmt.Println("Creating animal data to turn into json via marshal...")
	fmt.Println()

	a := newAnimal("Canine", 4, "Chichi", "Mixed")
	a2 := newAnimal("Canine", 4, "Bowie", "Black")

	l := []animal{a, a2}

	fmt.Println("Printing data of animals to show all visible and" +
		"non visible data")
	fmt.Println(l)
	fmt.Println()
	data, error := json.Marshal(l)

	if error != nil {
		log.Fatal(error)
	}

	fmt.Println("Json data:", string(data))
	fmt.Println()
	fmt.Println("Unmarshalling json data...")

	var animalList []animal
	var jsonAnimal animal
	var jsonAnimal2 animal

	fmt.Println(animalList)
	error = json.Unmarshal(data, &animalList)

	if error != nil {
		log.Fatal(error)
	}
	jsonAnimal = animalList[0]
	jsonAnimal2 = animalList[1]

	fmt.Println(animalList)
	fmt.Println("json animals:", jsonAnimal, "-", jsonAnimal2)
	fmt.Println()
	// Here we can see the externally visible (to json package) variables
	// have values and the rest (colour and VisibleButExcludedColour)
	// still have their nil value (as they were assigned their nil value
	// upon creation of the animal object).

	fmt.Println("category:", jsonAnimal.Category, "-",
		"num legs:", jsonAnimal.NumLegs, "-",
		"name:", jsonAnimal.Name, "-",
		"colour", jsonAnimal.colour, "-",
		"Colour", jsonAnimal.Colour, "-",
		"VisibleButExcludedColour", jsonAnimal.VisibleButExcludedColour, "-")

	fmt.Println()
	fmt.Println("Ending marshalling/unmarshalling example...")
	fmt.Println()
	fmt.Println("Starting encoding/decoding example...")
	fmt.Println()

	chichi := newAnimal("Canine", 4, "Chichi", "Mixed")
	fmt.Println(chichi) // show all content inside animal

	// json.NewEncoder(os.Stdout) returns a variable of type *Encoder
	// which writes to Stdout (console - prints to screen essentially).
	// As the function Encode() in the json package has a receiver of
	// type *Encoder (its a method of the type *Encoder), we can then
	// call .Encode on the value returned and provide some data to encode.
	error = json.NewEncoder(os.Stdout).Encode(chichi)
	if error != nil {
		log.Fatal(error)
	}

	data, error = json.Marshal(chichi)

	if error != nil {
		log.Fatal(error)
	}

	var chichi2 animal
	fmt.Println("Chichi2 after init:", chichi2)
	fmt.Println("Json string data:", string(data))

	reader := strings.NewReader(string(data))
	error = json.NewDecoder(reader).Decode(&chichi2)
	fmt.Println("Chichi2 after decode:", chichi2)
}
