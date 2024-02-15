/*
- https://www.w3schools.com/go/go_functions.php
- https://www.w3schools.com/go/go_function_parameters.php
- https://www.w3schools.com/go/go_function_returns.php
- https://www.w3schools.com/go/go_function_recursion.php
*/

package main

import (
	"fmt"
)

func functions_tutorial() {
	familyNameMain()
	familyNameMain2()
	fmt.Println(myFunction(1, 2))
	fmt.Println(myFunction2(1, 2))
	fmt.Println(myFunction3(1, 2))
	// Here, we store the return value in a variable called total
	total := myFunction4(1, 2)
	fmt.Println(total)
	// Here, myFunction() returns one integer (result) and one string (txt1)
	fmt.Println(myFunction5(5, "Hello"))
	// Here, we store the two return values into two variables (a and b)
	a, b := myFunction5(5, "Hello")
	fmt.Println(a, b)
	/* Here, we want to omit the first returned value (result - which is stored
	in variable c) */
	_, d := myFunction5(5, "Hello")
	fmt.Println(d)
	/* Here, we want to omit the second returned value (txt1 - which is stored
	in variable b) */
	c, _ := myFunction5(5, "Hello")
	fmt.Println(c)
}

/* Parameters and Arguments

Information can be passed to functions as a parameter. Parameters act as
variables inside the function.

Parameters and their types are specified after the function name, inside the
parentheses. You can add as many parameters as you want, just separate them with
a comma
*/

func familyName(fname string) {
	fmt.Println("Hello", fname, "Said")
}

func familyNameMain() {
	familyName("Liam")
	familyName("Jenny")
	familyName("Anja")
}

//------------------------------------------------------------------------------

/* Multiple Parameters

Inside the function, you can add as many parameters as you want
*/

func familyName2(fname string, age int) {
	fmt.Println("Hello", fname, "Said, who is", age, "years old")
}

func familyNameMain2() {
	familyName2("Liam", 30)
	familyName2("Jenny", 40)
	familyName2("Anja", 50)
}

//------------------------------------------------------------------------------

/* Function Returns

If you want the function to return a value, you need to define the data type of
the return value (such as int, string, etc), and also use the return keyword
inside the function
*/

/* Function Return Example

Here, myFunction() receives two integers (x and y) and returns their addition
(x + y) as integer (int)
*/

func myFunction(x int, y int) int {
	return x + y
}

//------------------------------------------------------------------------------

/* Named Return Values

In Go, you can name the return values of a function.

Here, we name the return value as result (of type int), and return the value
with a naked return (means that we use the return statement without specifying
the variable name)
*/

func myFunction2(x int, y int) (result int) {
	result = x + y
	return
}

/* The example above can also be written like this. Here, the return statement
specifies the variable name */

func myFunction3(x int, y int) (result int) {
	result = x + y
	return result
}

//------------------------------------------------------------------------------

/* Store the Return Value in a Variable

You can also store the return value in a variable, like this
*/

func myFunction4(x int, y int) (result int) {
	result = x + y
	return
}

//------------------------------------------------------------------------------

/* Multiple Return Values

Go functions can also return multiple values.

Here, myFunction() returns one integer (result) and one string (txt1)
*/

func myFunction5(x int, y string) (result int, txt1 string) {
	result = x + x
	txt1 = y + " World!"
	return
}
