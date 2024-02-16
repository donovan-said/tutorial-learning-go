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

	testcount(1)

	fmt.Println(factorial_recursion(4))
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

/*
Function Returns

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

/*
Named Return Values

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

//------------------------------------------------------------------------------

/* Recursion

Go accepts recursion functions. A function is recursive if it calls itself and
reaches a stop condition.

In the following example, testcount() is a function that calls itself. We use
the x variable as the data, which increments with 1 (x + 1) every time we
recurse. The recursion ends when the x variable equals to 11 (x == 11).
*/

func testcount(x int) int {
	if x == 11 {
		return 0
	}
	fmt.Println(x)
	return testcount(x + 1)
}

/*
Recursion is a common mathematical and programming concept. This has the benefit
of meaning that you can loop through data to reach a result.

The developer should be careful with recursion functions as it can be quite easy
to slip into writing a function which never terminates, or one that uses excess
amounts of memory or processor power. However, when written correctly recursion
can be a very efficient and mathematically-elegant approach to programming.

In the following example, factorial_recursion() is a function that calls itself.
We use the x variable as the data, which decrements (-1) every time we recurse.
The recursion ends when the condition is not greater than 0 (i.e. when it is 0).
*/

func factorial_recursion(x float64) (y float64) {
	if x > 0 {
		y = x * factorial_recursion(x-1)
	} else {
		y = 1
	}
	return
}
