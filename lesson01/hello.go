package main

import "fmt"

// const spanish = "Spanish"
// const french = "French"
// const englishHelloPrefix = "Hello, "
// const spanishHelloPrefix = "Hola, "
// const frenchHelloPrefix = "Bonjour, "
const (
	spanish = "Spanish"
	french = "French"
	englishHelloPrefix = "Hello, "
	spanishHelloPrefix = "Hola, "
	frenchHelloPrefix = "Bonjour, "
)

// func main() {
// 	fmt.Println("Hello, World!")
// }

func Hello() string {
	return englishHelloPrefix + "World!"
}

func HelloWithArgs(name string, language string) string {

	// If the name is empty, return a default greeting
	if name == "" {
		name = "World"
	}

	// if/else statements approach
	if (language == spanish) {
		return spanishHelloPrefix + name + "!"
	}

	if (language == french) {
		return frenchHelloPrefix + name + "!"
	}

	return greetingPrefix(language) + name + "!"
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
	return
}

func main() {
	fmt.Println(Hello())
	fmt.Println(HelloWithArgs("Chris", ""))
}