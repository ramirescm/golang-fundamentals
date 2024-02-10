package main

import "fmt"

func main() {
	for y := 33; y < 122; y++ {
		fmt.Printf("%d -> %v\n", int(y), string(y))
	}
}
