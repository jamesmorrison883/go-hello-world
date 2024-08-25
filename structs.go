package main

import "fmt"

func dataStructuresExamples() {

	//new struct called user
	type user struct {
		ID        int
		FirstName string
		LastName  string
	}

	//long form of declaring and initializing a struct. Here we create a variable called 'u' of type 'user'
	var u user
	u.ID = 1
	u.FirstName = "Arthur"
	u.LastName = "Dent"
	fmt.Println(u)

	//short form of declaring and initializing a struct
	u2 := user{ID: 1, FirstName: "Arthur", LastName: "Dent"}
	fmt.Println(u2)

	//short form of declaring and initializing a struct, demonstrating multi line struct initialization
	u3 := user{ID: 1,
		FirstName: "Arthur",
		LastName:  "Dent",
	}
	fmt.Println(u3)
}
