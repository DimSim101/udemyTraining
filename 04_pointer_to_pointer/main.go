package main

import "fmt"

func zero(z *int) {
	fmt.Println(z)
	*z = 0
}

func main() {
	x := 5
	y := 3

	// This is a ptr to y
	intPtr := &y
	fmt.Printf("Y is %v and intPtr is %v \n", y, *intPtr)

	*intPtr = 8
	fmt.Printf("Y is %v and intPtr is %v \n", y, *intPtr)

	// This is a ptr to a pointer - intPtr (a ptr to y). Hence it is a
	// pointer to a pointer to y.
	ptrintPtr := &intPtr
	fmt.Printf("ptrintPtr is %v \n", **ptrintPtr)

	**ptrintPtr = 2
	fmt.Printf("Y is %v and intPtr is %v and ptrintprt is %v \n", y,
		*intPtr, **ptrintPtr)

	fmt.Println(&x)
	zero(&x)
	fmt.Println(x) // x is 0

}
