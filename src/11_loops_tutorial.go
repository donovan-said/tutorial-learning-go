// https://www.w3schools.com/go/go_loops.php

package main

import (
	"fmt"
)

func loops_tutorial() {
	/*
		Go For Loops

		The for loop loops through a block of code a specified number of times.

		The for loop is the only loop available in Go.

		The for loop can take up to three statements

		- statement1 Initializes the loop counter value.
		- statement2 Evaluated for each loop iteration. If it evaluates to TRUE,
		  the loop continues. If it evaluates to FALSE, the loop ends.
		- statement3 Increases the loop counter value.

		Note: These statements don't need to be present as loops arguments.
		However, they need to be present in the code in some form.
	*/

	//--------------------------------------------------------------------------

	/*
		This example will print the numbers from 0 to 4

		i:=0; - Initialize the loop counter (i), and set the start value to 0
		i < 5; - Continue the loop as long as i is less than 5
		i++ - Increase the loop counter value by 1 for each iteration
	*/

	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	//--------------------------------------------------------------------------

	/*
		This example counts to 100 by tens:

		- i:=0; - Initialize the loop counter (i), and set the start value to 0
		- i <= 100; - Continue the loop as long as i is less than or equal to 100
		- i+=10 - Increase the loop counter value by 10 for each iteration
	*/

	for i := 0; i <= 100; i += 10 {
		fmt.Println(i)
	}

	//--------------------------------------------------------------------------

	/*
		The continue Statement

		The continue statement is used to skip one or more iterations in the
		loop. It then continues with the next iteration in the loop.

		This example skips the value of 3
	*/

	for i := 0; i < 5; i++ {
		if i == 3 {
			continue
		}
		fmt.Println(i)
	}

	//--------------------------------------------------------------------------

	/*
		The break Statement

		The break statement is used to break/terminate the loop execution.

		This example breaks out of the loop when i is equal to 3
	*/

	for i := 0; i < 5; i++ {
		if i == 3 {
			break
		}
		fmt.Println(i)
	}

	//--------------------------------------------------------------------------

	/*
		Nested Loops
		It is possible to place a loop inside another loop.

		Here, the "inner loop" will be executed one time for each iteration of
		the "outer loop"
	*/

	adj := [2]string{"big", "tasty"}
	fruits := [3]string{"apple", "orange", "banana"}
	for i := 0; i < len(adj); i++ {
		for j := 0; j < len(fruits); j++ {
			fmt.Println(adj[i], fruits[j])
		}
	}

	//--------------------------------------------------------------------------

	/*
		The Range Keyword

		The range keyword is used to more easily iterate over an array, slice or
		map. It returns both the index and the value.

		This example uses range to iterate over an array and print both the
		indexes and the values at each (idx stores the index, val stores the
		value):
	*/

	for idx, val := range fruits {
		fmt.Printf("%v\t%v\n", idx, val)
	}

	/*
		Tip: To only show the value or the index, you can omit the other output
		using an underscore (_).

		Here, we want to omit the indexes (idx stores the index, val stores the
		value)
	*/

	for _, val := range fruits {
		fmt.Printf("%v\n", val)
	}

	/*
		Here, we want to omit the values (idx stores the index, val stores the
		value)
	*/

	for idx, _ := range fruits {
		fmt.Printf("%v\n", idx)
	}
}
