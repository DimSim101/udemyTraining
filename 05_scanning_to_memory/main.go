package main

import "fmt"

const metersToYards float64 = 1.09361

func main() {
	// Note to self on how this robust Scan function works - I was
	// initially confused by how it magically processes passed in values
	// of types other than float64 (i.e. string 'a'), a value too large
	// (i.e. 9999999999...) etc. It also nicely allows for negative
	// numbers etc. If an error is returned and no bytes are scanned in,
	// then the value of the variable which was meant to hold the scanned
	// value will the equivalent of 0 for its type.

	// For example, in the example code below, we attempt to scan a
	// float64 value into var meters. If the scan is unsuccessful and an
	// error is thrown, no value is placed inside the meters variable.
	// Hence, it will still have its initial value, which we know in GO
	// due to the var meters float64 (static type) syntax, will have a
	// value of 0. This explains why when scanning valid float64 values,
	// the program works as expected, and if an error occurs (either
	// via being unable to convert type of scanned value to float64 or some
	// other reason), the value of meters will be what it is at
	// initialization (0), and hence the program will display output
	// accordingly. Below is a summary of the fmt.Scan process from
	// scanning to returning a value.

	// This fmt.Scan function is quite amazing it is given a pointer the
	// meters variable in order to know where to scan to. It then calls
	// FScan, which sets up a scan state and then proceeds to scan the
	// value in (via doScan). doScan is responsible for iterating over
	// the values scanned in and calling scanOne() on each. Within
	// scanOne is a switch statement that "derives the scanner from the
	// type of the argument", which in the case below derives the scanner
	// to be of type float64. Within the switch case for float64,
	// covertFloat is called which attempts to convert the values scanned
	// in into a float. Finally, this is where strconv.ParseFloat is
	// called on the scanned in value (in this case meters), which then
	// returns an error if the value is not of type float, is too large etc.
	// Obviously errors can occur at other points too, however this
	// explains the graceful handling of non float64 typed values being
	// passed in (i.e. strings).

	var meters float64
	fmt.Print("Enter meters swam: ")

	// grab out the number of bytes scanned (0 if fail, > 0 if success)
	// and the error if there is one
	scannedBytes, error := fmt.Scan(&meters)

	// Check for error due to incompatible type (attempting to scan in
	// something other than a value of type float64.
	if error != nil {
		fmt.Println(error)
	}

	fmt.Println("scanned in", scannedBytes, "bytes")
	fmt.Printf("meters scanned in value is: %v \n", meters)

	yards := meters * metersToYards
	fmt.Println(meters, "meters is", yards, "yards")
}
