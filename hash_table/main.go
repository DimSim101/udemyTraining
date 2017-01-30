package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
)

const numBuckets = 12

func main() {

	// This example is for a very basic hashtable implementation
	// using keys of type int and values of type string.

	stokerTracker := map[string]int{}

	result, error := http.Get("https://www.gutenberg.org/files/345/345.txt")
	if error != nil {
		log.Fatal(error)
	}

	scanner := bufio.NewScanner(result.Body)
	defer result.Body.Close()

	scanner.Split(bufio.ScanWords)
	buckets := make([][]string, numBuckets)

	fmt.Println("Buckets after initial make:", buckets)

	// create a slice for each bucket to hold all the strings
	for i := 0; i < numBuckets; i++ {
		buckets = append(buckets, []string{})
	}

	fmt.Println("Buckets after adding slice of strings to each slice:", buckets)

	for scanner.Scan() {
		// extract the word we have in the scanner buffer
		word := scanner.Text()
		// calculate which bucket to put it in
		n := hashBucket(word)
		// append the word to the bucket at index n (add it to the
		// slice of strings at index n of the buckets slice - buckets[n]).
		index := len(buckets[n])

		buckets[n] = append(buckets[n], word)

		// keep track of the index of each word in the bucket
		// so we can look up words directly later
		// In a "proper" hashtable we wouldn't do it like this
		// (obviously we cant use a hashtable when trying to create one
		// haha). Instead of just appending the word to the slice at
		// the bucket[n], we would precalculate the index for this
		// bucket too so that we could figure it out later. Otherwise
		// we can tell what bucket a word is in, but not at which index
		// inside that bucket.
		// For example, the Go docs mention using a hash function on
		// a value to generate a hash key (similar idea to HashBucket below
		// but more complex and better hash spread equality).
		// This hash function returns a 16 byte hash key. This 16 byte
		// value is then split into its upper and lower 8 byte chunks.
		// The lower 8 bytes are used to index the bucket array itself
		// (in this case which bucket slice to add to). The upper 8 bytes
		// are used to index the bucket slices themselves (which index
		// to place the value inside the specified bucket).
		// If we image HashBucket returned this kind of value, we could
		// do something like:

		// wordLocation := HashBucket("testing")
		// upperIndex, lowerIndex = split(wordLocation) - pseudocode not
		// real split function, split into 8 byte chunks.
		// value := buckets[lowerIndex][upperIndex]
		// This way for any value specified, we can rehash it and
		// get its location.
		// Otherwise, as in this basic implementation,
		// we cannot tell what index we have stored the value even
		// though it is there. In order to test if a value exists (as
		// an example), we would have to go through every entry in the
		// bucket and then return true if we found it or something similar).
		// With the precalulcation of the index for both buckets, we can
		// do these type of "exists" and similar lookups instantly.
		stokerTracker[word] = index
	}
	// fmt.Println(buckets[6][0:1]) // Print the first word in the slice of strings
	// for the 7th bucket (index 6). // Lol happened to be Bram so motivated
	// the making of stokerTracker.
	bramLocation := buckets[hashBucket("Bram")][stokerTracker["Bram"]]
	stokerLocation := buckets[hashBucket("Stoker")][stokerTracker["Stoker"]]
	fmt.Println(bramLocation, stokerLocation, ":)")
}

func hashBucket(word string) int {
	return int(word[0]) % numBuckets
}
