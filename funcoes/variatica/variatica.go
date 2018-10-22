package main

import (
	"fmt"
)

func main() {
	fmt.Printf("Média: %.2f \n", media(7.7, 8.3, 9.4, 5.6))

	//utilizando slice como parametro para uma função variatica
	aprovados := []string{"Maria", "Pedro", "Miguel", "Antonio"}
	imprimirAprovados(aprovados...)

}

func media(numeros ...float64) float64 {
	total := 0.0

	for _, valor := range numeros {
		total += valor
	}

	return total / float64(len(numeros))
}

func imprimirAprovados(aprovados ...string) {
	fmt.Println("Lista de aprovados:")
	for i, aprovado := range aprovados {
		fmt.Printf("%d) %s \n", i, aprovado)
	}
}
