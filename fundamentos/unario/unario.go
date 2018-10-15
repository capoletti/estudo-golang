package main

import (
	"fmt"
)

func main() {
	x := 1
	y := 2

	//go tem apenas as operações posfixadas
	x++
	y--

	fmt.Println("x:", x, " Y:", y)
}
