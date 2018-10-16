package main

import (
	"fmt"
)

func imprimirResultado(nota float64) {

	//if retira os parenteses da expressão, o bloco de chaves é obrigatório
	if nota >= 7 {
		fmt.Println("Aprovado com nota", nota)
	} else if nota >= 6 {
		fmt.Println("Recuperação com nota", nota)
	} else {
		fmt.Println("Reprovado com nota", nota)
	}
}

func main() {
	imprimirResultado(7.5)
	imprimirResultado(6.0)
	imprimirResultado(5.9)
}
