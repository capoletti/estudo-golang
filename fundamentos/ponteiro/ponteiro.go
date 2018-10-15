package main

import (
	"fmt"
)

//GO não possui funcoes aritméticas em ponteiro
func main() {
	i := 1

	var p *int

	p = &i //obtendo o endereço da variavel i para o ponteiro
	*p++   //alterando o valor pelo ponteiro
	i++

	fmt.Println(p, "==", &i, ":", *p, "==", i)

}
