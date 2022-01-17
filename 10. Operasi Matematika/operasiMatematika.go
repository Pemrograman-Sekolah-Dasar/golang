package main

import "fmt"

func main() {
	var a int = 80
	var b int = 20
	var c int = a + b
	var d int = a - b
	var e int = a * b
	var f int = a / b
	var g int = a % b

	fmt.Println(c) // 100
	fmt.Println(d) // 60
	fmt.Println(e) // 1600
	fmt.Println(f) // 4
	fmt.Println(g) // 0
}
