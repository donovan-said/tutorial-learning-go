// https://www.w3schools.com/go/go_struct.php

package main

import (
	"fmt"
)

/* Go Structures
A struct (short for structure) is used to create a collection of members of
different data types, into a single variable.

While arrays are used to store multiple values of the same data type into a
single variable, structs are used to store multiple values of different data
types into a single variable.

A struct can be useful for grouping data together to create records.
*/

type Person struct {
	name   string
	age    int
	job    string
	salary int
}

func struct_tutorial() {
	/*
		To access any member of a structure, use the dot operator (.) between
		the structure variable name and the structure member
	*/

	var person1 Person
	var person2 Person

	// Person 1 specifications
	person1.name = "Hege"
	person1.age = 45
	person1.job = "Teacher"
	person1.salary = 6000

	// Person 2 specifications
	person2.name = "John"
	person2.age = 24
	person2.job = "Maketing"
	person2.salary = 4500

	// Access and print Person 1 info
	// fmt.Println("Name: ", person1.name)
	// fmt.Println("Age: ", person1.age)
	// fmt.Println("Job: ", person1.job)
	// fmt.Println("Salary: ", person1.salary)

	// Access and print Person 2 info
	// fmt.Println("Name: ", person2.name)
	// fmt.Println("Age: ", person2.age)
	// fmt.Println("Job: ", person2.job)
	// fmt.Println("Salary: ", person2.salary)

	printPerson(person1)
	printPerson(person2)

}

func printPerson(pers Person) {
	fmt.Println("Name: ", pers.name)
	fmt.Println("Age: ", pers.age)
	fmt.Println("Job: ", pers.job)
	fmt.Println("Salary: ", pers.salary)
}
