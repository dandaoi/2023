package main

import (
	"fmt"
	"reflect"
	"strconv"
)

// var ( // Global variables
// 	/* name, course = "Nigel Pouton", "Getting started with Kubernetes"
// 	module, clip = 4, 2 */ // Best practice to leave variables separated for better understanding of the code
// 	name                   = "Nigel Pouton"
// 	course                 = "Getting startd with Kubernetes"
// 	module                 = "4"
// 	clip                   = 2
// )

func varFramework() { // if local variable is declared, it must be used.

	name := "Nigel Pouton"
	course := "Getting startd with Kubernetes"
	module := "4"
	clip := 2 // short assignment statement is used in local variables, because one can not use simply var = value without declaring the variable in the long form, if done so, the code will not read the variable correctly

	// var clip int
	// clip =2  ->  long format or
	// var clip int =2 -> long inline

	fmt.Println("Name and course are set to", name, "and", course, ".")
	fmt.Println("Module and clip are set to", module, "and", clip, ".")
	fmt.Println("Name is of type", reflect.TypeOf(name))
	fmt.Println("Module is of type", reflect.TypeOf(module))
	/*total := module + clip
	fmt.Println("Module plus clip equals", total)*/

	iModule, err := strconv.Atoi(module) // Atoi retorna 2 tipos, por conta disso ela popula 2 var, primeira com o valor em int zerado, segundo a msg de erro
	if err == nil {
		total := iModule + clip //Local variable, short assignment statement(:=)
		fmt.Println("Module plus clip equals", total)
	}

	fmt.Println("Memory address of  *course* variable is", &course)

	//var ptr *string = &course &->  points to the pointer address

	ptr := &course // forma de short assignment statement para ponteiro.
	fmt.Println("Pointing course variable at address", ptr, "which holds this value", *ptr)
	fmt.Println("Pointing course variable at address", &ptr)
	// Usando dessa forma, o segundo ponteiro referencia para o primeiro ponteiro, prevenindo que ambos fiquem utilizando a mesma informação, dessa forma, o segundo ponteiro simplesmente referencia o primeiro

}
