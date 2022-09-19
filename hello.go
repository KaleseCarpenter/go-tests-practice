//Writing Go with tests cycle
/* Write a test
Make the compiler pass
Run the test, see that it fails and check the error message is meaningful
Write enough code to make the test pass
Refactor */

//In Go, public functions start with a capital letter and private ones start with a lowercase.
//We don't want the internals of our algorithm to be exposed to the world, so we made this function private.

package main

import "fmt"

const spanish = "Spanish"
const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const french = "French"
const frenchHelloPrefix = "Bonjour, "
const name = "Madame Kalese"

func Hello(name string, language string) string {
	if name == "" {
		name = "world"
	}
	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case french:
		prefix = frenchHelloPrefix
	case spanish:
		prefix = spanishHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}

func main() {
	fmt.Println(Hello("world", ""))
	fmt.Println(frenchHelloPrefix + name)
}
