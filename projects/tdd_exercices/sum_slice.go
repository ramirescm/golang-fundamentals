package main

func SumSlice(numbers []int) int {
	sum := 0
	for _, value := range numbers {
		sum += value
	}

	return sum
}

// v2
func SumAll(numbersToSum ...[]int) []int {
	var total []int
	for _, numeros := range numbersToSum {
		total = append(total, SumSlice(numeros))
	}

	return total
}

// v1
// func SumAll(numbersToSum ...[]int) (sumNumbers []int) {
// 	totalNumbers := len(numbersToSum)
// 	sumNumbers = make([]int, totalNumbers)

// 	for i, numeros := range numbersToSum {
// 		sumNumbers[i] = SumSlice(numeros)
// 	}

// 	return
// }
