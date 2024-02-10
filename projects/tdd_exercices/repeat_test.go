package main

import "testing"

func TestRepeat(t *testing.T) {
	result := Repeat("a", 4)
	expected := "aaaa"

	if expected != result {
		t.Errorf("expected '%s', result '%s'", expected, result)
	}
}

func TestRepeatNTimes(t *testing.T) {
	result := Repeat("a", 5)
	expected := "aaaaa"

	if expected != result {
		t.Errorf("expected '%s', result '%s'", expected, result)
	}
}

func BenchmarkRepetir(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 4)
	}
}
