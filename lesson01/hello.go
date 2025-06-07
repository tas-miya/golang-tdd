package main

import "fmt"

const (
	spanish = "Spanish"
	french = "French"
	englishHelloPrefix = "Hello, "
	spanishHelloPrefix = "Hola, "
	frenchHelloPrefix = "Bonjour, "
)


func Hello() string {
	return englishHelloPrefix + "World!"
}

func HelloWithArgs(name string, language string) string {

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