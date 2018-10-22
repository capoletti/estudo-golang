package main

import (
	"fmt"
)

func main() {
	resultado, _ := fatorial(5)
	fmt.Println(resultado)

	_, err := fatorial(-2)
	fmt.Println(err)

	fmt.Println(fatorial2(5))
}

func fatorial(n int) (int, error) {
	switch {
	case n < 0:
		return -1, fmt.Errorf("número inválido: %d", n)
	case n == 0:
		return 1, nil
	default:
		fatorialAnterior, _ := fatorial(n - 1)
		return n * fatorialAnterior, nil
	}
}

func fatorial2(n uint) uint {
	switch {
	case n == 0:
		return 1
	default:
		return n * fatorial2(n-1)
	}
}
