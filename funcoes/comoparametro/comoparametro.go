package main

import (
	"fmt"
)

func main() {
	fmt.Println(exec(multiplicar, 2, 3))
}

func multiplicar(p1 int, p2 int) int {
	return p1 * p2
}

func exec(funcao func(int, int) int, p1, p2 int) int {
	return funcao(p1, p2)
}
