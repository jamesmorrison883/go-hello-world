package main

import (
	"fmt"
)

func arraysMapsExamples() {

	//declaring an initial array of 3 integers, print the output. Declared 'long form'
	var arr1 [3]int
	arr1[0] = 1
	arr1[1] = 2
	arr1[2] = 3
	fmt.Println(arr1)

	//declaring an initial array of 3 integers, print the output. Declared 'short form'
	arr2 := [3]int{1, 2, 3}
	fmt.Println(arr2)

	//demostrating the use of slices, which are a reference to an array. ':' represents a slice that includes all elements of an array
	slice := arr1[:]
	fmt.Println(slice)

	//demonstrating that slices are references to arrays, and that changing the value of an element in a slice, changes the value in the array
	arr1[1] = 42
	slice[2] = 27
	fmt.Println(arr1, slice)

	//demonstrating that slices can be created without an array, and that the slice is a reference to an array that is created behind the scenes
	slice2 := []int{1, 2, 3}
	fmt.Println(slice2)

	//demonstrating that slices can be appended to, and that the underlying array will be resized if necessary
	slice2 = append(slice2, 4, 42, 27)
	fmt.Println(slice2)

	//demonstrating that slices can be copied, and that the underlying array will be resized if necessary
	slice3 := make([]int, 2)
	copy(slice3, slice2)
	fmt.Println(slice3)

	//demonstrating that maps can be created, and that the key-value pairs can be accessed
	m := map[string]int{"foo": 42}
	fmt.Println(m)
	fmt.Println(m["foo"])

	//demonstrating that maps can be updated, and that the key-value pairs can be accessed
	m["foo"] = 27
	fmt.Println(m)

	//demonstrating that maps can be deleted, and that the key-value pairs can be accessed
	delete(m, "foo")
	fmt.Println(m)

}
