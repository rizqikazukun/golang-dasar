package main

import "fmt"

func main() {
	var nilaiAkhir byte = 90
	var absensi byte = 60

	var lulusAbsensi bool = absensi > 70
	var lulusNa bool = nilaiAkhir > 70
	var lulusSemua bool = lulusAbsensi && lulusNa

	fmt.Println(lulusAbsensi)
	fmt.Println(lulusNa)
	fmt.Println(lulusSemua)

}
