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

	/* The Println() function is similar to Print() with the difference that a
	whitespace is added between the arguments, and a newline is added at the
	end */

}
