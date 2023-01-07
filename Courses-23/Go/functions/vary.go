package main

//variadic functions - varying number of trading arguments passed on a function

import (
	"fmt"
)

func main() {

	bestFinish := championshipFinishes(12, 5, 6, 4, 3, 3)
	fmt.Println(bestFinish)
}

func championshipFinishes(finishes ...int) int { // ...int -> slice of ints, like a list of ints
	best := finishes[0]

	for _, i := range finishes {
		if i < best {
			best = i
		}
	}
	return best
}
