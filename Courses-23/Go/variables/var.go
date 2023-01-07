package main

import (
	"fmt"
	"os"
)

func main3() {
	name := os.Getenv("USERNAME")
	course := "Getting started with Kubernetes"
	fmt.Println("\nHi", name, "your current course is", course)
	updateCourse(&course)
	fmt.Println("You're currently watching", course)
	// updatedNameMessage := updateName(&name) --  this is using the value of return
	// fmt.Println(updatedNameMessage)
	// fmt.Println("Your name has been update to", name) -- this is not using the value of return, but using the updated statement on the variable.
}

// func updateName(name *string) string {
// 	*name = "Amanda Scaranello Tank"
// 	fmt.Println("Your name now is", *name)
// 	return *name + " is your new name :) " // might be useless, if not using it, just need to remove the type for the return of the func statement
// }

func updateCourse(course *string) string {

	*course = "Getting started with Docker"

	fmt.Println("Updated course to", *course)
	return *course
	// updates the copy of the variable not the variable itself, so value of course will not be updated.
}
