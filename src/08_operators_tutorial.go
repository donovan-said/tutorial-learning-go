/*
- https://www.w3schools.com/go/go_operators.php
- https://www.w3schools.com/go/go_arithmetic_operators.php
- https://www.w3schools.com/go/go_assignment_operators.php
- https://www.w3schools.com/go/go_comparison_operators.php
- https://www.w3schools.com/go/go_logical_operators.php
- https://www.w3schools.com/go/go_bitwise_operators.php
*/

package main

import (
	"fmt"
)

func operators_tutorial() {

	var a = 15 + 25
	fmt.Println(a)

	/*
		Although the + operator is often used to add together two values, it can
		also be used to add together a variable and a value, or a variable and
		another variable
	*/

	var (
		sum1 = 100 + 50    // 150 (100 + 50)
		sum2 = sum1 + 250  // 400 (150 + 250)
		sum3 = sum2 + sum2 // 800 (400 + 400)
	)
	fmt.Println(sum3)

	/*
		Assignment Operators

		Assignment operators are used to assign values to variables.

		In the example below, we use the assignment operator (=) to assign the
		value 10 to a variable called x
	*/

	var x = 10
	fmt.Println(x)

	// The addition assignment operator (+=) adds a value to a variable
	x += 5
	fmt.Println(x)

	/*
		Comparison Operators

		Comparison operators are used to compare two values.

		Note: The return value of a comparison is either true (1) or false (0).

		In the following example, we use the greater than operator (>) to find
		out if 5 is greater than 3
	*/

	var c = 5
	var d = 3
	fmt.Println(c > d) // returns 1 (true) because 5 is greater than 3
}
