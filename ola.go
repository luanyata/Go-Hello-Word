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

	return prefixSalutation(language) + name
}

func prefixSalutation(language string) (prefix string) {
	switch language {
	case localeFr:
		prefix = prefixFr
	case localeEs:
		prefix = prefixEs
	default:
		prefix = prefixPtBr
	}

	return
}

func main() {
	fmt.Println(Hello("mundo", ""))
}
