package main

import (
	"fmt"
)

func mainElements() {

	mySlice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(mySlice[3])

	mySlice[1] = 0
	fmt.Println(mySlice)

	sliceOfSlice := mySlice[2:5]
	fmt.Println(sliceOfSlice)
}
