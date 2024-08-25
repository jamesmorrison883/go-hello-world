package main

import "fmt"

func primitiveDataTypeExamples() {

	//explicit variable declaration and initialization
	var i int
	i = 42
	fmt.Println(i)

	//still explicit example, but more concise and demonstrates that float32 is a primitive type in golang
	var f float32 = 3.14
	fmt.Println(f)

	// implicit variable declaration, and initialization
	firstName := "Arthur"
	fmt.Println(firstName)

	// implicit variable declaration, that is concise, demonstrating bool is primitive to golang
	b := true
	fmt.Println(b)

	// implicit variable declaration, that is concise, demonstrating complex nubmers type is primitive to golang
	c := complex(3, 4)
	fmt.Println(c)

	// demonstrates yet another way of interacting with complex numbers, while also illustrating
	r, im := real(c), imag(c)
	fmt.Println(r, im)
}
