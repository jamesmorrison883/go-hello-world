package main

func pointerExamples() {
	// start with a simple (long form) declaration of a variable, i.
	var x = 42
	println("value of x:", x)

	// declare a pointer variable, ptr, that is uninitialized.
	var ptr *int

	// Initialize ptr by assigning the address of i to it, using the 'addressof' operator
	ptr = &x

	//print the value held at the memory address stored in ptr, using the derefrence operator
	println("value of *ptr:", *ptr)

	// change the value of i by changing the value at the address stored in ptr, using the derefrence operator
	*ptr = 21

	//print the value of ptr, which is an address.
	println("value of ptr:", ptr)

	//print the value of i, which has been changed by changing the value at the address stored in p
	println("value of x:", x)
}
