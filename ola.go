package main

import "fmt"

const prefixHelloPtBr = "Olá "

func Hello(name string) string {
	return prefixHelloPtBr + name
}

func main() {
	fmt.Println(Hello("Mundo"))
}
