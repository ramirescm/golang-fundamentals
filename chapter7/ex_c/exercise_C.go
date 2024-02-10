package main

import (
	"fmt"
	"time"
)

func main() {
	bornDate := 1988
	currentYear := time.Now().Year()
	for bornDate <= currentYear {
		fmt.Println(bornDate)
		bornDate++
	}

	fmt.Println("------------")
	bornDate = 1988
	currentYear = time.Now().Year()
	for {
		if currentYear >= bornDate {
			fmt.Println(currentYear)
			currentYear--
		} else {
			break
		}

	}
}
