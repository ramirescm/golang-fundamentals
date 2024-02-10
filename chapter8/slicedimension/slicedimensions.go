package main

import "fmt"

func main() {
	ss := [][]int{
		[]int{1, 2, 3}, // index 0
		[]int{4, 5, 6}, // index 1
		[]int{7, 8, 9}, // index 2
	}

	fmt.Println(ss[0][1])
}
