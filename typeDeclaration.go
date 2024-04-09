package main

import (
	"fmt"
	"strconv"
)

func main() {
	type NoKTP string // type alias

	var no NoKTP = NoKTP("11111111")
	fmt.Println(no)
	fmt.Println(strconv.Itoa(111111))

}
