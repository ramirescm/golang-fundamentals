package main

import (
	"fmt"
	"testing"
)

func TestSum(t *testing.T) {
	result := Add(2, 2)
	expected := 4

	if result != expected {
		t.Errorf("expected '%d', result '%d'", expected, result)
	}
}

func ExAdd() {
	result := Add(1, 5)
	fmt.Println(result)
	// Output: 6
}
