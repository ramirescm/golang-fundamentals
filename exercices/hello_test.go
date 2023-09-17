package main

import (
	"testing"
)

func TestHello(t *testing.T) {
	checkRigthMessage := func(t *testing.T, result, expected string) {
		t.Helper()
		if result != expected {
			t.Errorf("result '%s', expected '%s'", result, expected)
		}
	}

	t.Run("says hello to people", func(t *testing.T) {
		result := Hello("Jarvis", "")
		expected := "Hello, Jarvis"
		checkRigthMessage(t, result, expected)
	})

	t.Run("says 'Hello, world' when an empty string is passed", func(t *testing.T) {
		result := Hello("", "")
		expected := "Hello, world"
		checkRigthMessage(t, result, expected)
	})

	t.Run("greatting in Spanish", func(t *testing.T) {
		result := Hello("Elodie", "spanish")
		expected := "Hola, Elodie"
		checkRigthMessage(t, result, expected)
	})

	t.Run("greatting in French", func(t *testing.T) {
		result := Hello("Maria", "french")
		expected := "Bonjour, Maria"
		checkRigthMessage(t, result, expected)
	})
}
