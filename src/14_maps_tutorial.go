// https://www.w3schools.com/go/go_maps.php

package main

import (
	"fmt"
)

/*
Go Maps
Maps are used to store data values in key:value pairs.

Each element in a map is a key:value pair.

A map is an unordered and changeable collection that does not allow duplicates.

The length of a map is the number of its elements. You can find it using the
len() function.

The default value of a map is nil.

Maps hold references to an underlying hash table.

Go has multiple ways for creating maps.
*/

func maps_tutorial() {

	//--------------------------------------------------------------------------

	/*
		Create Maps Using var and :=

		This example shows how to create maps in Go. Notice the order in the
		code and in the output */

	var a = map[string]string{"brand": "Ford", "model": "Mustang", "year": "1964"}
	b := map[string]int{"Oslo": 1, "Bergen": 2, "Trondheim": 3, "Stavanger": 4}

	fmt.Printf("a\t%v\n", a)
	fmt.Printf("b\t%v\n", b)

	//--------------------------------------------------------------------------

	/*
		Create Maps Using make()Function

		This example shows how to create maps in Go using the make()function.
	*/

	var c = make(map[string]string) // The map is empty now
	c["brand"] = "Ford"
	c["model"] = "Mustang"
	c["year"] = "1964"

	d := make(map[string]int)
	d["Oslo"] = 1
	d["Bergen"] = 2
	d["Trondheim"] = 3
	d["Stavanger"] = 4

	fmt.Printf("c\t%v\n", c)
	fmt.Printf("d\t%v\n", d)

	//--------------------------------------------------------------------------

	/*
		Create an Empty Map

		There are two ways to create an empty map. One is by using the make()
		function and the other is by using the following syntax.

		var a map[KeyType]ValueType

		Note: The make()function is the right way to create an empty map. If you
		make an empty map in a different way and write to it, it will causes a
		runtime panic.

		This example shows the difference between declaring an empty map using
		with the make() function and without it.
	*/

	var e = make(map[string]string)
	var g map[string]string

	fmt.Println(e == nil)
	fmt.Println(g == nil)

	//--------------------------------------------------------------------------

	/*
		Access Map Elements

		You can access map elements by

		value = map_name[key]
	*/

	fmt.Println(a["brand"])

	//--------------------------------------------------------------------------

	/*
		Update and Add Map Elements

		Updating or adding an elements are done by:

		map_name[key] = value

		This example shows how to update and add elements to a map.
	*/

	fmt.Println(a)

	a["year"] = "1970" // Updating an element
	a["color"] = "red" // Adding an element

	fmt.Println(a)

	//--------------------------------------------------------------------------

	/*
		Remove Element from Map

		Removing elements is done using the delete() function.

		delete(map_name, key)
	*/

	delete(a, "year")

	fmt.Println(a)

	//--------------------------------------------------------------------------

	/*
		Check For Specific Elements in a Map

		You can check if a certain key exists in a map using:

		val, ok :=map_name[key]

		If you only want to check the existence of a certain key, you can use
		the blank identifier (_) in place of val.
	*/

	var h = map[string]string{"brand": "Ford", "model": "Mustang", "year": "1964", "day": ""}

	val1, ok1 := h["brand"]
	val2, ok2 := h["color"]
	val3, ok3 := h["day"]
	_, ok4 := h["model"]

	fmt.Println(val1, ok1)
	fmt.Println(val2, ok2)
	fmt.Println(val3, ok3)
	fmt.Println(ok4)

	/*
		Example Explained
		In this example, we checked for existence of different keys in the map.

		The key "color" does not exist in the map. So the value is an empty
		string ('').

		The ok2 variable is used to find out if the key exist or not. Because we
		would have got the same value if the value of the "color" key was empty.
		This is the case for val3.
	*/

	//--------------------------------------------------------------------------

	/*
		Maps Are References

		Maps are references to hash tables.

		If two map variables refer to the same hash table, changing the content
		of one variable affect the content of the other.
	*/

	var i = map[string]string{"brand": "Ford", "model": "Mustang", "year": "1964"}
	j := i

	fmt.Println(i)
	fmt.Println(j)

	j["year"] = "1970"
	fmt.Println("After change to j:")

	fmt.Println(i)
	fmt.Println(j)

	//--------------------------------------------------------------------------

	/*
		Iterate Over Maps

		You can use range to iterate over maps.

		This example shows how to iterate over the elements in a map. Note the
		order of the elements in the output.
	*/

	l := map[string]int{"one": 1, "two": 2, "three": 3, "four": 4}

	for k, v := range l {
		fmt.Printf("%v : %v, ", k, v)
	}

	//--------------------------------------------------------------------------

	/*
		Iterate Over Maps in a Specific Order

		Maps are unordered data structures. If you need to iterate over a map in
		a specific order, you must have a separate data structure that specifies
		that order.
	*/
}
