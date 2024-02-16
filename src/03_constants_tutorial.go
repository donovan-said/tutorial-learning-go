// https://www.w3schools.com/go/go_constants.php

package main

import (
	"fmt"
)

const PI = 3.14

func constants_tutorial() {

	fmt.Println(PI)

	/*
		There are two types of constants:

		- Typed constants
		- Untyped constants
	*/

	// Typed constants are declared with a defined type
	const A int = 1
	fmt.Println(A)

	// Untyped constants are declared without a type
	const B = 1
	fmt.Println(B)

	// Multiple constants can be grouped together into a block for readability
	const (
		EG1 int = 1
		EG2     = 3.14
		EG3     = "Hi!"
	)

	fmt.Println(EG1)
	fmt.Println(EG2)
	fmt.Println(EG3)
}
