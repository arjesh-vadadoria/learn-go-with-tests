package main

import (
	"fmt"
)

const (
	spanish = "Spanish"
	french  = "French"

	spanishHelloPrefix = "Hola, "
	frenchHelloPrefix  = "Bonjour, "
	englishHelloPrefix = "Hello, "

	suffix = "!"
)

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	return greetingPrefix(language) + name + suffix
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case spanish:
		prefix = spanishHelloPrefix
	case french:
		prefix = frenchHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return prefix
}

func main() {
	fmt.Println(Hello("Arjesh", "Spanish"))
}
