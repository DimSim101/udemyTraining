package main

import (
	"fmt"
	"github.com/DimSim101/udemyTraining/03_visibility_and_ptrs/vis"
)

func main() {
	fmt.Printf("Initial vis yourName value: %v \n", *vis.GetNamePtr())
	var name = vis.GetNamePtr()
	fmt.Printf("Initial local yourName value: %v \n", *name)
	*name = "testing my name"
	fmt.Printf("LOCAL IS %v \n", *name)
	fmt.Printf("NON LOCAL IS %v \n", *vis.GetNamePtr())

	fmt.Println()

	fmt.Printf("Initial MyName is %v \n", vis.MyName)
	var ourName = &vis.MyName
	fmt.Printf("Initial ourName is %v \n", *ourName)
	*ourName = "BLAH"
	fmt.Printf("After MyName is %v \n", vis.MyName)
	fmt.Printf("After ourName is %v \n", *ourName)

	fmt.Println()

	vis.PrintVar()
}
