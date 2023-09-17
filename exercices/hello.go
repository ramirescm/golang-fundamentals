package main

import (
	"fmt"
)

const spanish = "spanish"
const french = "french"

const greetingsEnglish = "Hello, "
const greetingsSpanish = "Hola, "
const greetingsFrench = "Bonjour, "

func Hello(name string, language string) string {
	if name == "" {
		name = "world"
	}

	return greattingsPrefix(language) + name
}

func greattingsPrefix(language string) (prefix string) {
	switch language {
	case french:
		prefix = greetingsFrench
	case spanish:
		prefix = greetingsSpanish
	default:
		prefix = greetingsEnglish
	}

	return
}

func main() {
	fmt.Println(Hello("Jarvis", "french"))
}
