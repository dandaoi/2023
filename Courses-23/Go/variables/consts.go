package main

import (
	"fmt"
)

func main1() {

	const c = 186000
	fmt.Println("The speed of light is", c, "miles per second")

	const km = c * 1.56
	fmt.Println("The speed of light is", km, "km per second")
	// 	c = 91039129032
	// 	fmt.Println("You cannot change constant values", c)
}
