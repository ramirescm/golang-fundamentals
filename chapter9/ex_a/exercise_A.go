package main

import "fmt"

var x [5]int

func main() {
	trees := [5]int{1, 2, 3, 4, 5}
	x[0] = 1

	fmt.Println(trees)
	fmt.Println(x)

	fmt.Printf("%T", trees)
}
