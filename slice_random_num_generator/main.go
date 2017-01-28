package main

import (
	"fmt"
	"io"
	"crypto/rand"
)

func main() {
	var nonce [16]byte
	fmt.Println(nonce)

	// Since nonce is of type [16]byte (an array of length 16 of type byte) and
	// io.ReadFull takes a reader and a slice as its args we cant do: io.ReadFull(rand.Reader, nonce)
	// Instead, we can convert the array nonce into a slice by using slice expressions and
	// grabbing all its values through default values [:]

	io.ReadFull(rand.Reader, nonce[:])
	fmt.Println(nonce)
}
