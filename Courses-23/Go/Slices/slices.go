package main

// Arrays have a fixed size, a slice can be resized and are built on top of arrays(slices of an array) and no actual data stored in them, like pointers pointing to the actual data
// Slices are passed to functions by reference
import (
	"fmt"
)

func mainSlices() {

	courses := []string{"Docker & Kubernetes: The Big Picture",
		"Getting Started with Docker",
		"Getting Started with Kubernetes"}

	fmt.Printf("Length of slice is %d and capacity is %d \n",
		len(courses), cap(courses))

	for _, i := range courses {
		fmt.Println(i)
	}
}
