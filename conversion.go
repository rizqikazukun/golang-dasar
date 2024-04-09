package main

import "fmt"

func main() {
	var nilai32 int32 = 32768
	var nilai64 int64 = int64(nilai32)
	var nilai16 int16 = int16(nilai32)

	fmt.Println(nilai32) // initial value
	fmt.Println(nilai64) // same value with different type
	fmt.Println(nilai16) // cant change type and make value wrong

	var name = "Rizqi"
	var e = name[1]
	var eString = string(e)

	fmt.Println(name)
	fmt.Println(e)
	fmt.Println(eString)

}
