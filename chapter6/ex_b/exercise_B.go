package main

import "fmt"

func main() {
	if x := 10; x >= 10 {
		fmt.Println(x)
	}

	if x := 10; x > 10 {
		fmt.Println(x)
	} else if x < 10 {
		fmt.Println(x)
	} else {
		fmt.Println("igual")
	}
}
