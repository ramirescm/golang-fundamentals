package main

import "fmt"

// package level scope
var total int
var name = "Beans"
var quantity = 10
var isTruth = true
var price = 12.23

var a int
var b string
var c bool
var d float64

func main() {
	fmt.Println("hello world")
	fmt.Printf("%v %T\n", name, name)
	printProduct()
	printTypes()
	typeString()
}

func typeString() {
	fmt.Println("\n>> string literal and string interpreted <<")
	msgRaw := `\t dont interprete backslash \n test`
	msgInterpreted := "\t hello \n world"
	fmt.Println(msgRaw)
	fmt.Println(msgInterpreted)
}

func printTypes() {
	fmt.Println("\n>> types zero <<")
	fmt.Printf("%v - %T\n", a, a)
	fmt.Printf("%v - %T\n", b, b)
	fmt.Printf("%v - %T\n", c, c)
	fmt.Printf("%v - %T\n", d, d)
}

func printProduct() {
	fmt.Println("\n>> product <<")
	fmt.Println(name)
	fmt.Println(quantity)
	fmt.Println(isTruth)
	fmt.Println(price)
	fmt.Println(total)
}
