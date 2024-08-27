package main

import "fmt"

func main() {

	fmt.Println("Hello from a module, World!")
	primitiveDataTypeExamples()
	pointerExamples()
	arraysMapsExamples()
	dataStructuresExamples()
	port := 3000
	retries := 3
	// '_' is a write only variable, used to ignore the error return value
	_, err := startWebServer(port, retries)
	fmt.Println(port, err)

}
