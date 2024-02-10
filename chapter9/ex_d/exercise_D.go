package main

import "fmt"

func main() {
	x := make([]string, 26, 26)
	x = []string{"asf", "as"}

	fmt.Println(x, len(x), cap(x))

	for i := 0; i < len(x); i++ {
		fmt.Println(x[i])
	}
}
