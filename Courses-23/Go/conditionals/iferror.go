package main

import (
	"fmt"
	"os"
)

func mainIfError() {
	_, err := os.Open("./test.txt")

	if err != nil {
		fmt.Println("This is the error code:", err)
	}

	fmt.Println("This is the error code:", err)

}
