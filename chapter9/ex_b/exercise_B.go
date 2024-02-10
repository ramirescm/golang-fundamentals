package main

import "fmt"

func main() {
	trees := [5]int{1, 2, 3, 4, 5}

	for index, val := range trees {
		fmt.Println(index, val)
	}

}
