package main

import "fmt"

func main() {

	friends := map[string]int{
		"joosh": 12345,
		"maria": 5678,
	}

	fmt.Println(friends)

	fmt.Println(friends["maria"])

	friends["jjj"] = 4444

	fmt.Println(friends)

	fmt.Println(friends["inexistente"])

	// coma ok idiom - ok true tem, false nao tem
	val, ok := friends["inexistente"]
	fmt.Println(val, ok)
}
