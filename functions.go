package main

import (
	"fmt"
)

//func 'functionname' (input parameters) (returntype1, returntype2, n) {}
//idiomatic go tends to rely on return types that are errors

func startWebServer(port int, retries int) (int, error) {
	fmt.Println("Starting server...")
	//do important things
	fmt.Println("Server started on port", port)
	fmt.Println("Number of retries", retries)
	return port, nil

}
