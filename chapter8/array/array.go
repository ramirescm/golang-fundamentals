package main

import "fmt"

var x [5]int

func main() {
	x[0] = 1
	x[1] = 12
	x[2] = 133

	fmt.Println(x[0])
	fmt.Println(x[1])
	fmt.Println(x[2])
}
