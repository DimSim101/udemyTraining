package main

import "fmt"

func main() {

	b := true

	// interesting that using this type of initialization is accessible
	// within the entire if/else chunk, but not outside of it
	// and this differs to initializing a variable inside the if block
	// (cannot be accessed within the else if / else etc. - see below).
	if food := "Chocolate"; !b {
		blah := "bbb"
		fmt.Println("\n" + food + "\n")
		fmt.Println(blah)
	} else if b {
		// cannot access blah here (as it is within the closure of the
		// if statement curly braces)
		// blah = "asd"
		fmt.Println("not in same section, though init variable"+
			"is still accessible as we can see ... ", food, "\n")
	} else {
		// food variable can still be accessed as we are still inside
		// the "if" block (still within scope)
		food = "blah"
		fmt.Println(food)
		// same applies to blah above here, outside the closure of if {}
		//blah = "asd"
	}

	// no longer in the scope to access this variable - line below will fail
	// food = "abc"

	// You would think this would work but it doesn't (even though its the
	// same logically. If and switch statements actually have an optional
	// initialization part (which must then end with a semi colon). Hence,
	// why the below method doesn't work (the init and the comparison are not
	// part of the same section of the statement)

	/*
		if food := "Chocolate" && b {
			fmt.Println(food)
		}
	*/

	// Just like the if statement, initialization within the switch also
	// requires a semi colon, followed by the name of the value to switch on
	// (in this case myvar)

	// Similarly, the variable myvar (initialized outside the switch block)
	// is accessible anywhere within the switch.
	switch myvar := "bbb"; myvar {
	// testing string comparison with "" and ``
	case `bbb`:
		// variables created within the switch block are specific to
		// the case they are initialized in
		// This is the exact same as the if case shown above -
		// initializing a variable inside the initialization statement
		// does not limit it to a specific block within the statement
		// (if or switch) but instead limits the variables scope to the
		// entire statements block!
		blah := "test"
		fmt.Println("We found it! :D Switch init works. Just like with " +
			"if statements if we initialize a variable in the " +
			"statement, we must have a semi colon after it, followed " +
			"by the usual statement (i.e. comparison when in an if, " +
			"statement to switch cases on in switch)" + "\n")
		fmt.Println(blah)
		fallthrough
	case "aaa":
		// the line below will fail as blah is not defined
		// blah = "aaa"
		fmt.Println("this should never happen...unless we fallthrough hehe :) \n")
		// threw this in as fallthrough is not allowed in type switch's
		fallthrough
	default:
		fmt.Println("default... can do something that is required in " +
			"all the switch cases above that have fallthrough, " +
			"or just something that is required for default behaviour \n")

		//myvar accessible within the block of the switch statement as expected
		fmt.Println("SWITCHED ON VALUE:", myvar)
	}

	// once outside the switch block, we cannot access myvar anymore
	// myvar = "ccc" - This line will fail

}
