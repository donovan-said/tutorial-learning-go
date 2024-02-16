/*
- https://www.w3schools.com/go/go_output.php
- https://www.w3schools.com/go/go_formatting_verbs.php
*/

package main

import (
	"fmt"
)

func output_tutorial() {
	var i, j string = "Hello", "World"

	// The Print() function prints its arguments with their default format
	fmt.Print(i)
	fmt.Print(j)

	// If we want to print the arguments in new lines, we need to use \n.
	fmt.Print(i, "\n")
	fmt.Print(j, "\n")

	/* It is also possible to only use one Print() for printing multiple
	variables. */

	fmt.Print("\n", i, "\n", j)

	// If we want to add a space between string arguments, we need to use " "
	fmt.Print("\n", i, " ", j, "\n")

	// Print() inserts a space between the arguments if neither are strings
	var x, y = 10, 20

	fmt.Print(x, y)

	/*
		The Println() function is similar to Print() with the difference that a
		whitespace is added between the arguments, and a newline is added at the
		end
	*/

	fmt.Println(i, j)

	/*
		The Printf() function first formats its argument based on the given
		formatting verb and then prints them

		- %v is used to print the value of the arguments
		- %T is used to print the type of the arguments
	*/

	fmt.Printf("i has value: %v and type: %T\n", i, i)
	fmt.Printf("j has value: %v and type: %T", j, j)
}
