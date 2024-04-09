package main

import "fmt"

func main() {
	var names [2]string // golang array is length restricted

	names[0] = "rizqi" // set array value
	names[1] = "pratama"

	fmt.Println(names[0])
	fmt.Println(names[1])

	fmt.Println(len(names)) // get array length

	var values = [3]int{
		2,
		3,
	}

	// this is how directly set array values
	// and if array value is integer, the default value is 0

	fmt.Println(values)

	var cars = [...]string{
		"audi",
		"toyota",
	} // this is how write array without length restriction

	fmt.Println(cars)

}
