package main

import (
	"fmt"
	"github.com/DimSim101/udemyTraining/03_visibility_and_ptrs/vis"
)

type family struct {
	mum person
	me person
	related bool
}

type person struct {
	first string
	last string
	age int
}

// This struct contains a single embedded type (person) which can then
// be accessed directly i.e. DoubleZero.first to get first name of person.
type doubleZero struct {
	person
	LicenseToKill bool
}

// This struct was an attempt at holding multiple embedded types. I guess we can,
// but not multiple embedded types of the SAME TYPE. I.e. I cannot have multiple
// embedded variables of type person inside a single struct.
// These anonymous variables (as they have no name through which we can refer to
// them) are "promoted" which means they are accessible in the context of the
// struct itself rather than needing to "step down a level" into the person
// type.
type tripleZero struct {
	declaredPerson person
	person
	//person - This throws an error as we have duplicate names, even with
	// the optional tag included.
	toBeOrNotToBe bool
}


// this is a method that is a function connected to some type
// in this case, a function connected to the person type.
// The receiver essentially connects the function to the type of the receiver
// This way any variable of type person has access to the function fullName().
func (p person) fullName() string {
	return p.first + " " + p.last
}

// variable and method overriding can exist in structs due to their encapsulation
// You can access variables and their defined methods via each level of the struct
// using its . syntax. See below for examples.
// This function will be called when using the family.fullName() syntax.
// If we want to call the underlying persons function, we must do family.person.fullName().
func (f family) fullName() string {
	// we used a dash here to differentiate the two elements
	return f.me.first + "-" + f.me.last
}

func main() {
	me := person{"David", "Aaron", 25}
	fmt.Println(me)
	fmt.Println(me.first, me.last, me.age)
	fmt.Println(me.fullName())

	// Both initialization techniques for a struct result in a struct
	// that has the zero values for all the variables it contains.

	// initialize struct using var default nil value -- same as below
	var p2 person
	fmt.Println(p2)
	fmt.Println(p2.age)

	// initialize struct to empty value -- same as above
	p3 := person{}
	fmt.Println(p3)
	fmt.Println(p3.age)
	fmt.Println()

	// Pointer to p1 is just a pointer to the first variable in the struct
	// Presumably this is just the start of the memory chunk in which
	// all the struct variables are placed
	fmt.Printf("%p \n", &me)
	fmt.Printf("%p - %p - %p \n\n", &me.first, &me.last, &me.age)

	fmt.Printf("%p \n", &p2)
	fmt.Printf("%p - %p - %p \n\n", &p2.first, &p2.last, &p2.age)

	fmt.Printf("%p \n", &p3)
	fmt.Printf("%p - %p - %p \n", &p3.first, &p3.last, &p3.age)

	// visibility rules apply to naming structs just like any other variable
	// capital first letter = visible outside package.
	test := vis.VisiblePerson{}
	fmt.Println(test)
	fmt.Println()

	var nilFam family
	fmt.Println(nilFam)

	// We dont require the name: value syntax as below. This is just a
	// syntax element for clarity. See below for example without
	myFam := family{
		mum: person{"Mandy", "Aaron", 56},
		me: me,
		related: true,
	}

	fmt.Println(myFam)

	uglyFam := family{
		person{"Ugly", "Fugly", 69},
		person{"Who is who", "Here", 69},
		false,
	}

	fmt.Println(uglyFam)

	fmt.Println(myFam.mum)
	fmt.Println(uglyFam.mum)

	fmt.Println()
	fmt.Println("fullName method associated with person variable:", me.fullName())

	fmt.Println("Full name from family fullName() method:", myFam.fullName())
	fmt.Println("Full name from person fullName() method inside struct:", myFam.me.fullName())
	fmt.Println()

	d1 := doubleZero{
		person: person{
			first: "James",
			last: "Bond",
			age: 42,
		},
		LicenseToKill: true,
	}

	// as this is an embedded type, we can access its values without
	// requiring d1.person.blah.
	fmt.Println(d1.first, d1.last, d1.age, d1.LicenseToKill)
	fmt.Println()

	t1 := tripleZero{
		declaredPerson: person{
			first: "Shake",
			last: "Spear",
			age: 99,
		},
		person: person{
			first: "Test",
			last: "Second",
			age: 100,
		},
		toBeOrNotToBe: true,
	}

	// This shows the difference between an embedded type and a non embedded type
	fmt.Println(t1.first, t1.last, t1.age, t1.toBeOrNotToBe)
	fmt.Println(t1.declaredPerson.first, t1.declaredPerson.last, t1.declaredPerson.age, t1.toBeOrNotToBe)
	fmt.Println()

	// Just had to check this was the same as the pointer examples above when
	// grabbing the address of the struct created.
	lastPerson := &person{"Last", "Person", 7}
	fmt.Printf("%p: %p - %p - %p \n", lastPerson, &lastPerson.first, &lastPerson.last, &lastPerson.age)
}
