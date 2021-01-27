package main

import (
	"fmt"
)

// iterateOverSlice is a fucntion to iterate given slice
func iterateOverSlice(data []int) {
	for i, val := range data {
		fmt.Println(i, val)
	}

}

func main() {
	fmt.Println("Introduction to slice operations")
	A := []int{1, 2, 3, 4, 5}
	iterateOverSlice(A)

}
