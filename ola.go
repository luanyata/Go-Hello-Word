package main

import "fmt"

const prefixHelloPtBr = "Olá "

func Hello(name string) string {
	if name == "" {
		name = "mundo"
	}
	return prefixHelloPtBr + name
}

func main() {
	fmt.Println(Hello("mundo"))
}
