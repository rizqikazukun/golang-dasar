package main

import "fmt"

func variable() {
	var name string

	name = "Rizqi Pratama"
	fmt.Println(name)

	name = "Rizqikazukun"
	fmt.Println(name) // reassign value

	var fname = "rizqi" // golang can know the type
	fmt.Println(fname)

	lname := "pratama" // var keyword is optional
	fmt.Println(lname)

	var (
		sFname = "Rizqi"
		sLname = "Pratama"
	) // Multiple variable declaration
	fmt.Println(sFname)
	fmt.Println(sLname)

}
