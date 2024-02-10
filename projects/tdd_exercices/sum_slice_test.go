package main

import (
	"reflect"
	"testing"
)

func TestSumSlice(t *testing.T) {
	t.Run("collection with 5 numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}
		result := SumSlice(numbers)
		expected := 15

		if expected != result {
			t.Errorf("result %d, expected %d, values %v", result, expected, numbers)
		}
	})

	t.Run("collection any size", func(t *testing.T) {
		numbers := []int{1, 2, 3}
		result := SumSlice(numbers)
		expected := 6

		if expected != result {
			t.Errorf("result %d, expected %d, values %v", result, expected, numbers)
		}
	})

	t.Run("sum everthing", func(t *testing.T) {
		result := SumAll([]int{1, 2}, []int{0, 9})
		expected := []int{3, 9}

		if !reflect.DeepEqual(result, expected) {
			t.Errorf("result %v expected %v", result, expected)
		}
	})

}
