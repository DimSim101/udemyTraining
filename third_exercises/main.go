package main

import (
	"fmt"
	"sort"
)

type people []string
type peopleTwo []string
// These are not needed as we can use sort.StringSlice since we are
// sorting a slice of strings. However, I included them in the final example
// to show their use.
func (p people) Len() int { return len(p) }
func (p people) Swap(i,j int)  {
	// Because this is an expression, the right hand side is evaluated
	// before the left hand side.
	// This results in the RHS creating a tuple with the value (p[j], p[i])
	// Where p[j] and p[i] are their UNSWAPPED VALUES.
	// Following this, the left hand side is evaluated (left hand side
	// variables are assigned their respective values) i.e. p[i] now = p[j].
	// Thus, the variables now have swapped their values without a temp variable
	p[i], p[j] = p[j], p[i]
}
func (p people) Less(i,j int) bool { return p[i] < p[j] }

// Here I am showing how you can override a types String() method which is
// called to return the strings value when printing. If we change the receiver
// to type people, all the cases where we used variables of type people below
// will have their printed values changed.
func (p peopleTwo) String() string {
	var output string
	for index, value := range p {
		output += fmt.Sprintf("STRINGY - %d : %s\n", index, value)
	}
	return output
}

func main()  {
	fmt.Println()
	fmt.Println("Example 1:")
	studyGroup := people{"Zeno", "John", "Al", "Jenny"}
	fmt.Println(studyGroup)

	stringSlice := sort.StringSlice(studyGroup)
	sort.Sort(stringSlice)
	fmt.Println(studyGroup)

	// all sort.Reverse does below is take an Interface and change its
	// Less() function to be the other way around i.e. p[j] < p[i]
	sort.Sort(sort.Reverse(stringSlice))
	fmt.Println(studyGroup)
	fmt.Println()

	fmt.Println("Example 2:")
	s := []string{"Zeno", "John", "Al", "Jenny"}
	fmt.Println(s)

	stringSlice = sort.StringSlice(s)
	sort.Sort(stringSlice)
	fmt.Println(s)
	sort.Sort(sort.Reverse(stringSlice))
	fmt.Println(s)
	// There is also a Sort() function attached to the StringSlice type
	// which just calls sort.Sort(). It is only there for convenience but
	// I figured I may aswell show its use here.
	sort.StringSlice(s).Sort()
	fmt.Println(s)
	fmt.Println()

	fmt.Println("Example 3:")
	n := []int{7, 4, 8, 2, 9, 19, 12, 32, 3}
	fmt.Println(n)

	intSlice := sort.IntSlice(n)
	sort.Sort(intSlice)
	fmt.Println(n)
	sort.Sort(sort.Reverse(intSlice))
	fmt.Println(n)
	fmt.Println()

	fmt.Println("My own test:")
	lastTest := people{"David", "Cassi", "Mandy", "Mark", "Oma", "Robbie"}
	fmt.Println(lastTest)
	lastTestInterface := people(lastTest)
	sort.Sort(lastTestInterface)
	fmt.Println(lastTest)
	sort.Sort(sort.Reverse(lastTestInterface))
	fmt.Println(lastTest)

	// Just like the strings, ints has its own "quick access" sort
	n2 := []int{79, 42, 18, 2, 9, 190, 12, 32, 3}
	sort.Ints(n2)
	fmt.Println(n2)
	fmt.Println()

	peopleTwoGroup := peopleTwo{"PeopleTwo", "Zeno", "John", "Al", "Jenny"}
	fmt.Print(peopleTwoGroup)
}
