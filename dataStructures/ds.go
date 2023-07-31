package datastructures

import (
	"fmt"
)

func CreateArray() {
	var array []int

	println("Capacity of Empty Array:", cap(array))

	for i := 0; i < 5; i++ {
		array = append(array, i)
	}

	fmt.Printf("array: %v\n", array)

	// Get the length of the slice
	length := len(array)
	println("Length:", length)

	//Capacity
	println("Capacity:", cap(array))

	//copy array
	tempArray := make([]int, len(array))
	copy(tempArray, array)
	fmt.Printf("temp Array: %v\n", tempArray)

	println("temp Array Cap:", cap(tempArray))

}
