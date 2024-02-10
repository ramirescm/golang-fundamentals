package main

import "fmt"

// atencao com alteracao de arrays subjacentes, ou altere o mesmo array ou copie com for

func main() {
	slice := []string{"banana", "maça", "jaca"}

	for index, value := range slice {
		fmt.Println(index, value)
	}

	fmt.Printf("\n\n")

	fmt.Println("add new element")
	slice = append(slice, "pessego")
	for index, value := range slice {
		fmt.Println(index, value)
	}

	fmt.Printf("\n\n")

	flavors := []string{"chocolate", "vanilla", "mellon", "lemon", "strawberry"}

	takeFirst := flavors[:1] // flavors[:1] ou flavors[0:1]
	fmt.Println(takeFirst)

	takeTwo := flavors[1:3]
	fmt.Println(takeTwo)

	takeLast := flavors[4:] // flavors[4:] ou flavors[0:1]
	fmt.Println(takeLast)

	fmt.Printf("\n\n all flavors \n\n")
	fmt.Println(flavors[:])
	fmt.Println(flavors)

	fmt.Printf("\n\nprint all flavors ")

	for i := 0; i < len(flavors); i++ {
		fmt.Println(flavors[i])
	}

	fmt.Printf("\n\n remove and print flavors ")

	removed := append(flavors[:2], flavors[4:]...)
	fmt.Println(removed)

	fmt.Printf("\n\n make pizzas \n\n ")

	pizza := make([]int, 4, 8)
	pizza[0], pizza[1], pizza[2], pizza[3] = 1, 2, 3, 4

	pizza = append(pizza, 5)
	pizza = append(pizza, 6)
	pizza = append(pizza, 7)
	pizza = append(pizza, 8)
	// element, size, capacity
	fmt.Println(pizza, len(pizza), cap(pizza))

	pizza = append(pizza, 8)
	// element, size, capacity
	fmt.Println(pizza, len(pizza), cap(pizza))
}
