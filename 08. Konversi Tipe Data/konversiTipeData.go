package main

import "fmt"

func main() {
	var nilai32 int32 = 1000
	var nilai64 int64 = int64(nilai32)
	fmt.Println(nilai64)

	var nama string = "Prasetiyo"
	var e byte = nama[0]
	var eString string = string(e)
	fmt.Println(eString)
}
