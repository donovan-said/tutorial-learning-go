// https://www.w3schools.com/go/go_arrays.php

package main

import "fmt"

func arrays_tutorial() {
	/*
		In Go, there are two ways to declare an array
		- With the var keyword
		- With the := sign
	*/

	var arr1 = [3]int{1, 2, 3}
	arr2 := [5]int{4, 5, 6, 7, 8}
	arr3 := [...]int{9, 10, 11}
	arr4 := []int{12, 13, 14}

	fmt.Println(arr1)
	fmt.Println(arr2)
	fmt.Println(arr3)
	fmt.Println(arr4)

	var cars = [4]string{"Volvo", "BMW", "Ford", "Mazda"}
	fmt.Println(cars)

	prices := [3]int{10, 20, 30}

	fmt.Println(prices[0])
	fmt.Println(prices[1])
	fmt.Println(prices[2])

	/*
		You can also change the value of a specific array element by referring
		to the index number.

		This example shows how to change the value of the third element in the
		prices array
	*/

	prices[2] = 50
	fmt.Println(prices)

	/*
		If an array or one of its elements has not been initialized in the code,
		it is assigned the default value of its type.

		The default value for int is 0, and the default value for string is ""
	*/

	arr5 := [5]int{}              //not initialized
	arr6 := [5]int{1, 2}          //partially initialized
	arr7 := [5]int{1, 2, 3, 4, 5} //fully initialized

	fmt.Println(arr5)
	fmt.Println(arr6)
	fmt.Println(arr7)

	/*
		It is possible to initialize only specific elements in an array.

		This example initializes only the second and third elements of the array
		- 1:10 means: assign 10 to array index 1 (second element).
		- 2:40 means: assign 40 to array index 2 (third element).
	*/

	arr8 := [5]int{1: 10, 2: 40}
	fmt.Println(arr8)

	// The len() function is used to find the length of an array

	fmt.Println(len(arr7))
	fmt.Println(len(arr8))
}
