package main

import "fmt"

func slice() {
	cars := [...]string{"audi", "toyoya", "Mercy", "BMW", "VW", "Tesla"}

	slice := cars[:]
	fmt.Println(slice)

	fmt.Println(len(slice))
	fmt.Println(cap(slice))

	topCar := cars[:2]
	fmt.Println(len(topCar))
	fmt.Println(cap(topCar))

	var motocycle []string = make([]string, 2, 2)

	motocycle[0] = "honda"
	// motocycle2 :=

	fmt.Println(append(motocycle, "yamaha"))
	fmt.Println(motocycle)
	// fmt.Println(len(motocycle2))

}
