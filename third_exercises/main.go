package main

import (
	"fmt"
	"sort"
)

type people []string

// These are not needed as we can use sort.StringSlice since we are
// sorting a slice of strings. However, I included them in the final example
// to show their use.
func (p people) Len() int { return len(p) }
func (p people) Swap(i,j int)  { p[i], p[j] = p[j], p[i] }
func (p people) Less(i,j int) bool { return p[i] < p[j] }

func main()  {
	fmt.Println()
	fmt.Println("Example 1:")
	studyGroup := people{"Zeno", "John", "Al", "Jenny"}
	fmt.Println(studyGroup)

	stringSlice := sort.StringSlice(studyGroup)
	sort.Sort(stringSlice)
	fmt.Println(studyGroup)

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
}
