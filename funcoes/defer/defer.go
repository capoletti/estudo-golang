package main

import (
	"fmt"
)

func main() {
	fmt.Println(obterValor(5001))
	fmt.Println(obterValor(4999))
}

func obterValor(numero int) int {
	defer fmt.Println("fim!")

	if numero >= 5000 {
		fmt.Println("valor alto...")
		return 5000
	}

	fmt.Println("valor baixo...")
	return numero
}
