package main

import (
	"fmt"
)

const spanish = "Spanish"
const french = "French"
const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjour, "

// Named return value allows you to show exactly what is returned up front
// This function is private to this package because it starts with a lowercase character
func greetingPrefix(language string) (prefix string) {
	switch language {
	case french:
		prefix = frenchHelloPrefix
	case spanish:
		prefix = spanishHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	// You can return a named return value with just `return`
	return
}

func Hello(name, language string) string {
	if name == "" {
		name = "World"
	}
	return greetingPrefix(language) + name
}

func main() {
	fmt.Println(Hello("", ""))
}

// Cycle:
// 1. Write a test
// 2. Make the compiler pass
// 3. Run the test, see that it fails and check that the error message is meaningful
// 4. Write enough code to make the test pass
// 5. Refactor
