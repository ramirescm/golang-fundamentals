package main

import "fmt"

func main() {
	x := 100
	fmt.Printf("%d\t%b\t%#x\n", x, x, x)
	xy := x << 1
	fmt.Printf("%d\t%b\t%#x\n", xy, xy, xy)
	xx := xy >> 1
	fmt.Printf("%d\t%b\t%#x\n", xx, xx, xx)
}
