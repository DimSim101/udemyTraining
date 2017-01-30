package main

import (
	"fmt"
)

func main() {

	s := make([]string, 3, 9)

	fmt.Println(s)

	// any value left out below defaults to len(s) - be careful

	fmt.Println(s[:])
	fmt.Println(s[0:])

	// Possible out of range error if:
	// low > len (i.e. low > len && < cap) and the other side is left blank (then defaults to len which is invalid) for example:

	//fmt.Println(s[8:]) // This becomes s[8:3] (illegal) after replacement as any BLANK is replaced with len(s). This is now invalid (cant go negative direction)
	fmt.Println(s[:8]) // This becomes s[3:8] (legal) after replacement. We can do stuff like this because here the placement of len(s) as the default doesn't break the syntax.

	fmt.Println(s[6:9]) // As mentioned, we can extract values > len(s) and <= cap(s) using slicing expressions.
	// Remember we couldnt do fmt.Println(s[6:]) as this becomes fmt.Println(s[6:len(s)] which becomes fmt.Println(s[6:3]) which is invalid and throws out of bounds.

	// append will just add to the end rather than replacing the default " " values inside that are set when creating using make
	s = append(s, "blah")
	fmt.Println(s)          // all of s (the slice)
	fmt.Println(s[:])       // all of s from 0 to len(s) -- same as above
	fmt.Println(s[:cap(s)]) // all of s and its underlying array - not same

	// If we want to replace the default created empty string items in the slice, we must refer to their index. Can loop over from i; i < len(s); i ++ and change them all if we wanted.
	s[0] = "one"
	s[1] = "two"
	s[2] = "three"
	fmt.Println(s)
	fmt.Println(s[:cap(s)]) // all of s and its underlying array - not same
}
