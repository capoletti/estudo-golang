package main

import (
	"fmt"
)

func main() {
	n := 1

	inc1(n) //por valor
	fmt.Println(n)

	//& recupera o endeço de uma variável
	inc2(&n) //por referencia
	fmt.Println(n)
}

func inc1(n int) {
	n++
}

//ponteiro é representado por *
func inc2(n *int) {
	//* é utilizao para desreferenciar
	*n++
}
