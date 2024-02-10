package main

import "fmt"

func main() {
	people := [][]string{
		[]string{
			"aa", "cs", "dd",
		},
		[]string{
			"ae", "cd", "dd",
		},
		[]string{
			"as", "cf", "df",
		},
	}

	for i, v := range people {
		fmt.Println(i, v)
	}

	// Declare an empty slice of slices
	fruits := [][]string{}

	// Populate the slice in another line using append
	person1 := []string{"Alice", "maca"}
	fruits = append(fruits, person1)

	person2 := []string{"Bob", "banana"}
	fruits = append(fruits, person2)

	person3 := []string{"Charlie", "kiwi"}
	fruits = append(fruits, person3)

	// Access and print the data
	for _, person := range fruits {
		fmt.Printf("Name: %s, Fruit: %s\n", person[0], person[1])
	}
}
