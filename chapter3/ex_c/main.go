package chapter3

import "fmt"

var x int = 42
var y string = "James Bondd"
var z bool = true

func main() {

	s := fmt.Sprintf("%v\t%v\t%v\t", x, y, z)
	fmt.Print(s)
}
