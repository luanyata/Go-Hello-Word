package main

import "fmt"

const localeEs = "es"
const localeFr = "fr"
const prefixPtBr = "Olá "
const prefixEs = "Hola "
const prefixFr = "Bonjour "

func Hello(name string, language string) string {
	if name == "" {
		name = "mundo"
	}

	if language == localeEs {
		return prefixEs + name
	}

	if language == localeFr {
		return prefixFr + name
	}

	return prefixPtBr + name
}

func main() {
	fmt.Println(Hello("mundo", ""))
}
