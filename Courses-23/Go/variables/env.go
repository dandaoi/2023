package main

import (
	"fmt"
	"os"
)

func main2() {

	for _, env := range os.Environ() {
		fmt.Println(env)
	}

}
