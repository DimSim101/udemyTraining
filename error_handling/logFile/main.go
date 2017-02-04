package main

import (
	"os"
	"log"
	"fmt"
)

var logFile *os.File

func init() {
	// This will CREATE the file if it doesnt exist, else it will open the file
	// with WRITE ONLY (since were only writing logs) permissions and append to it.
	logFile, error := os.OpenFile("log.txt", os.O_WRONLY | os.O_CREATE | os.O_APPEND, 0666)
	// ORIGINAL LINE FROM TODD's EXAMPLE TO CREATE LOG FILE --
	// os.Create("log.txt") // create simple log txt file. Will overwrite
	// old file if it existed.
	if error != nil {
		// If we cant setup the log file, log the error to stdout and exit.
		log.Fatal(error)
	}

	log.SetOutput(logFile) // Set logging output to write to the log file.
	fmt.Println("Log output set to:", logFile.Name())
}

func main() {
	defer logFile.Close() // close the log file at the very end.

	file, error := os.Open("non-existant-file.txt") // attempt to open fake file.
	if error != nil {
		log.Println("Error:", error) // Log the error to the log file.
		// If we needed this file to open in order to do some stuff,
		// we could panic or fatal here.
		//log.Panic("Error:", error) // Will call exit after logging and
		// print the stack and exit status
		//log.Fatal("Error:", error) // Will call exit after logging and
		// print the exit status
	} else {
		// If we got here, we must've had a file called
		// non-existant-file.txt in the directory. Wtf lol?
		fmt.Println("File", file, "successfully opened")
	}

}
