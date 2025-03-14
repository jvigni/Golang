package main

import (
	"fmt"
)

func main() {
	// remove 4 from array
	array1 := []int{1, 2, 3, 4, 5, 6, 7}
	array2 := array1[:3]
	array3 := array1[4:]
	array4 := append(array2, array3...) // ... to convert array to numbers
	fmt.Println(array4)

}
