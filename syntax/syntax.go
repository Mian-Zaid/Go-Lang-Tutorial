package syntax

import "fmt"

var name string //declare variable

var id int = 10 //init variable

const pi = 3.14 //const

func print() {
	size := 10 //short var declare

	fmt.Println("size:", size)
}

func conditionalStatements() {
	var condition = 10
	if condition > 10 {
		// code
	} else if condition < 10 {
		// code
	} else {
		// code
	}
}
