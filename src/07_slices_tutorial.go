/*
- https://www.w3schools.com/go/go_slices.php
- https://www.w3schools.com/go/go_slices_modify.php
*/

package main

import (
	"fmt"
)

func slices_tutorial() {

	myslice1 := []int{}
	fmt.Println(len(myslice1))
	fmt.Println(cap(myslice1))
	fmt.Println(myslice1)

	myslice2 := []string{"Go", "Slices", "Are", "Powerful"}
	fmt.Println(len(myslice2))
	fmt.Println(cap(myslice2))
	fmt.Println(myslice2)

	// You can create a slice by slicing an array

	arr1 := [6]int{10, 11, 12, 13, 14, 15}
	myslice3 := arr1[2:4]

	fmt.Printf("myslice3 = %v\n", myslice3)
	fmt.Printf("lenght = %d\n", len(myslice3))
	fmt.Printf("capacity = %d\n", cap(myslice3))

	// The make() function can also be used to create a slice

	myslice4 := make([]int, 5, 10)
	fmt.Printf("myslice4 = %v\n", myslice4)
	fmt.Printf("length = %d\n", len(myslice4))
	fmt.Printf("capacity = %d\n", cap(myslice4))

	// with omiited capacity
	myslice5 := make([]int, 5)
	fmt.Printf("myslice5 = %v\n", myslice5)
	fmt.Printf("length = %d\n", len(myslice5))
	fmt.Printf("capacity = %d\n", cap(myslice5))

	/*
		You can access a specific slice element by referring to the index
		number.

		In Go, indexes start at 0. That means that [0] is the first element, [1]
		is the second element, etc
	*/

	prices := []int{10, 20, 30}

	fmt.Println(prices[0])
	fmt.Println(prices[2])

	// This example shows how to change the third element in the prices slice

	prices[2] = 50
	fmt.Println(prices[0])
	fmt.Println(prices[2])

	// This example shows how to append elements to the end of a slice

	prices = append(prices, 20, 21)
	fmt.Printf("prices = %v\n", prices)
	fmt.Printf("length = %d\n", len(prices))
	fmt.Printf("capacity = %d\n", cap(prices))

	/*
		To append all the elements of one slice to another slice, use the
		append() function

		This example shows how to append one slice to another slice
	*/

	myslice7 := []int{1, 2, 3}
	myslice8 := []int{4, 5, 6}
	myslice9 := append(myslice7, myslice8...)

	fmt.Printf("myslice9 = %v\n", myslice9)
	fmt.Printf("length = %d\n", len(myslice9))
	fmt.Printf("capacity = %d\n", cap(myslice9))

	/*
		Unlike arrays, it is possible to change the length of a slice

		This example shows how to change the length of a slice
	*/

	arr2 := [6]int{9, 10, 11, 12, 13, 16} // An array
	myslice10 := arr2[1:5]                // Slice array
	fmt.Printf("myslice10 = %v\n", myslice10)
	fmt.Printf("length = %d\n", len(myslice10))
	fmt.Printf("capacity = %d\n", cap(myslice10))

	myslice11 := arr2[1:3] // Change the length by re-slicing the array
	fmt.Printf("myslice11 = %v\n", myslice11)
	fmt.Printf("length = %d\n", len(myslice11))
	fmt.Printf("capacity = %d\n", cap(myslice11))

	myslice12 := append(myslice11, 20, 21, 22, 23) // Change length by appending items
	fmt.Printf("myslice12 = %v\n", myslice12)
	fmt.Printf("length = %d\n", len(myslice12))
	fmt.Printf("capacity = %d\n", cap(myslice12))

	/*
		Memory Efficiency

		When using slices, Go loads all the underlying elements into the memory.

		If the array is large and you need only a few elements, it is better to
		copy those elements using the copy() function.

		The copy() function creates a new underlying array with only the
		required elements for the slice. This will reduce the memory used for
		the program.

		This example shows how to use the copy() function
	*/

	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
	// Original slice
	fmt.Printf("numbers = %v\n", numbers)
	fmt.Printf("length = %d\n", len(numbers))
	fmt.Printf("capacity = %d\n", cap(numbers))

	// Create copy with only needed numbers
	neededNumbers := numbers[:len(numbers)-10]
	numbersCopy := make([]int, len(neededNumbers))
	copy(numbersCopy, neededNumbers)

	fmt.Printf("numbersCopy = %v\n", numbersCopy)
	fmt.Printf("length = %d\n", len(numbersCopy))
	fmt.Printf("capacity = %d\n", cap(numbersCopy))

}
