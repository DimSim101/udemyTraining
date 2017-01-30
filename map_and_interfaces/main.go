package main

import (
	"fmt"
	"net/http"
	"log"
	"io/ioutil"
)

type myReader struct {
	data int
}

// An attempt to create a Reader to see if it would just return these values.
// Just gets called in some sort of loop and runs out of memory
// Presumably we need to provide some sort of EOF to stop.
func (r myReader) Read(p []byte) (n int, err error)  {
	fmt.Println("read called")
	return 0, nil
}

func main() {

	result, error := http.Get("http://www-01.sil.org/linguistics/wordlists/english/wordlist/wordsEn.txt")
	if error != nil {
		log.Fatal(error)
	}

	data, error := ioutil.ReadAll(result.Body)
	fmt.Println(string(data))

	if error != nil {
		log.Fatal(error)
	}

	var r myReader
	fmt.Println(r)
	// data, error = ioutil.ReadAll(r) //This loops until out of memory
}
