/*
	- https://www.w3schools.com/go/go_variables.php
	- https://www.w3schools.com/go/go_variable_multi.php
	- https://www.w3schools.com/go/go_variable_naming_rules.php
*/

package main

import (
	"fmt"
)

/*
var
- Can be used inside and outside of functions
- Variable declaration and value assignment can be done separately

:=
  - Can only be used inside functions
  - Variable declaration and value assignment cannot be done separately (must be
    done in the same line)
*/
var global1 int
var global2 = 2

func variables_tutorial() {

	// If not set, the type of the value is inferred
	var student1 string = "John"
	var student2 = "Jane"
	x := 2

	fmt.Println(student1)
	fmt.Println(student2)
	fmt.Println(x)

	/* Variables without values initial value will be set to default value of
	it's type */
	var a string
	var b int
	var c bool

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)

	// Value assignment after declaration
	var person string
	person = "John"
	fmt.Println(person)

	// Global vars
	fmt.Println(global1)
	fmt.Println(global2)

	// It is possible to declare multiple variables in the same line
	var e, f, g, h int = 1, 3, 5, 7

	fmt.Println(e)
	fmt.Println(f)
	fmt.Println(g)
	fmt.Println(h)

	/* If the type keyword is not specified, you can declare different types of
	variables in the same line */

	var eg1, eg2 = 6, "Hello"
	eg3, eg4 := 7, "World"

	fmt.Println(eg1)
	fmt.Println(eg2)
	fmt.Println(eg3)
	fmt.Println(eg4)

	/* Multiple variable declarations can also be grouped together into a block
	for greater readability */

	var (
		eg5 int
		eg6 int    = 1
		eg7 string = "hello"
	)

	fmt.Println(eg5)
	fmt.Println(eg6)
	fmt.Println(eg7)

	// Camel Case
	myVariableName := "John"

	// Pascal Case
	MyVariableName := "John"

	// Snake Case
	my_variable_name := "John"

	fmt.Println(myVariableName)
	fmt.Println(MyVariableName)
	fmt.Println(my_variable_name)
}
