/*
- https://www.w3schools.com/go/go_switch.php
- https://www.w3schools.com/go/go_switch_multi.php
*/

package main

import (
	"fmt"
)

func switch_tutorial() {
	/*
		Use the switch statement to select one of many code blocks to be
		executed.

		The switch statement in Go is similar to the ones in C, C++, Java,
		JavaScript, and PHP. The difference is that it only runs the matched
		case so it does not need a break statement

		This is how it works:

		- The expression is evaluated once
		- The value of the switch expression is compared with the values of
		  each case
		- If there is a match, the associated block of code is executed
		- The default keyword is optional. It specifies some code to run if
		  there is no case match
	*/

	day := 4

	switch day {
	case 1:
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuesday")
	case 3:
		fmt.Println("Wednesday")
	case 4:
		fmt.Println("Thursday")
	case 5:
		fmt.Println("Friday")
	case 6:
		fmt.Println("Saturday")
	case 7:
		fmt.Println("Sunday")
	}

	/*
		It is possible to have multiple values for each case in the switch
		statement

		The example below uses the weekday number to return different text
	*/

	switch day {
	case 1, 3, 5:
		fmt.Println("Odd weekday")
	case 2, 4:
		fmt.Println("Even weekday")
	case 6, 7:
		fmt.Println("Weekend")
	default:
		fmt.Println("Invalid day of day number")
	}
}
