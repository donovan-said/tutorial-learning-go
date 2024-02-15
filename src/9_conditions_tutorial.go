/*
- https://www.w3schools.com/go/go_conditions.php
- https://www.w3schools.com/go/go_if_statement.php
- https://www.w3schools.com/go/go_else_statement.php
- https://www.w3schools.com/go/go_elseif_statement.php
- https://www.w3schools.com/go/go_nested_if.php
*/

package main

import (
	"fmt"
)

func conditions_tutorial() {
	/* In the example below, we test two values to find out if 20 is greater
	than 18. If the condition is true, print some text
	*/

	if 20 > 18 {
		fmt.Println("20 is greater than 18")
	}

	// We can also test variables

	x := 20
	y := 18
	if x > y {
		fmt.Println("x is greater than y")
	}

	/* Go if statement

	In this example, time (20) is greater than 18, so the if condition is
	false. Because of this, we move on to the else condition and print to the
	screen "Good evening". If the time was less than 18, the program would print
	"Good day"
	*/

	time := 20
	if time < 18 {
		fmt.Println("Good day.")
	} else {
		fmt.Println("Good evening.")
	}

	/* In this example, the temperature is 14 so the condition for if is false
	so the code line inside the else statement is executed
	*/

	temperature := 14
	if temperature > 15 {
		fmt.Println("It is warm out there")
	} else {
		fmt.Println("It is cold out there")
	}

	// N.b. Having the else brackets in a different line will raise an error

	/* Go if else Statement

	Use the else if statement to specify a new condition if the first condition
	is false */

	time_new := 22
	if time_new < 10 {
		fmt.Println("Goog morning.")
	} else if time < 20 {
		fmt.Println("Good day.")
	} else {
		fmt.Println("Good evening")
	}

	// Another example for the use of else if

	a := 14
	b := 14
	if a < b {
		fmt.Println("a is less than b")
	} else if a > b {
		fmt.Println("a is more than b")
	} else {
		fmt.Println("a and b are equal")
	}

	/* Note: If condition1 and condition2 are BOTH true, only the code for
	condition1 are executed */

	c := 30
	if c >= 10 {
		fmt.Println("x is larger than or equal to 10")
	} else if c > 10 {
		fmt.Println("x is larger than 20")
	} else {
		fmt.Println("x is less than 10")
	}

	/* Go Nested if Statement

	You can have if statements inside if statements, this is called a nested if.

	This example shows how to use nested if statements */

	num := 20
	if num >= 10 {
		fmt.Println("Num is more than 10")
		if num > 15 {
			fmt.Println("Num is also more than 15")
		}
	} else {
		fmt.Println("Num is less tna 10")
	}
}
