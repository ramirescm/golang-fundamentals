package main

import "fmt"

func main() {
	x := 0
	for {

		fmt.Println("x ", x)
		x++
		if x == 3 {
			break
		}
	}

	for y := 2; y < 4; y++ {
		fmt.Println("y ", y)
	}

	w := 5
	for w <= 25 {
		fmt.Println("w ", w)
		w = w * 2
	}
}
