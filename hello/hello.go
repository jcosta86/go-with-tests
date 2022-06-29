package main

import "fmt"

const (
	englishHelloPrefix = "Hello, "
	portuguesePrefix   = "Ola, "
	frenchPrefix       = "Bonjour, "
	portuguese         = "Portuguese"
	french             = "french"
)

func Hello(name, language string) string {
	if name == "" {
		name = "World"
	}
	prefix := greetingPrefix(language)
	switch language {
	case portuguese:
		prefix = portuguesePrefix
	case french:
		prefix = frenchPrefix
	}

	return prefix + name
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case portuguese:
		prefix = portuguesePrefix
	case french:
		prefix = frenchPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}

func main() {
	fmt.Println(Hello("Jeff", ""))
}
