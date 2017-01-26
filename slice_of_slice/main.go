package main

import (
"fmt"
"strings"
)

func main() {
	var records [][]string
	// student 1
	student1 := make([]string, 4)
	student1[0] = "Foster"
	student1[1] = "Nathan"
	student1[2] = "100.00"
	student1[3] = "74.00"
	// store the record
	records = append(records, student1)
	// student 2
	student2 := make([]string, 4)
	student2[0] = "Gomez"
	student2[1] = "Lisa"
	student2[2] = "92.00"
	student2[3] = "96.00"
	// store the record
	records = append(records, student2)

	myTestStudent := []string {"David", "Aaron", "100.00", "99.00", "98.00", "420.00"}
	records = append(records, myTestStudent)

	// print
	printTable(records)
}

// Tried to use an interface as the args and extract the values / switch on type
// rather than pass in a slice of slice of strings [][]string.
func printTable(x interface{}) {
	fmt.Println(strings.Repeat("-", 120))
	records, ok := x.([][]string)

	if ok {
		for i, record := range records {
			fmt.Printf("| %v: | ", i)
			for j := 0; j < len(record); j++ {
				lengthDiff := 15 - len(record[j])
				if lengthDiff > 0 {
					fmt.Print(strings.Repeat(" ", lengthDiff))
				}
				fmt.Printf(" %v |", record[j])
			}
			fmt.Println()
		}
	}

	fmt.Println(strings.Repeat("-", 120))
}
