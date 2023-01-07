package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	switch tmpNum := random(); tmpNum {
	case 0, 2, 4, 6, 8:
		fmt.Println("We got the following even number -", tmpNum)
	case 1, 3, 5, 7, 9:
		fmt.Println("We got the followin odd number -", tmpNum)
		// case "kubernetes":
		// 	fmt.Println("Case 1. kubernetes with lower case \"k\".")
		// case "Kubernetes":
		// 	fmt.Println("Case 2. Kubernetes with capital \"K\".")
		// case "K8S":
		// 	fmt.Println("Case 3. Kubernetes short name \"Kates\".")
		// case "Istio":
		// 	fmt.Println("Case 1. Service mesh is the future \"k\".")
		// default:
		// 	fmt.Println("Hit the default")
	}

	//fallthrough ->  will run the pertaining code block and fall to the next case lvl a
	//break ->
}

func random() int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(10)
}
