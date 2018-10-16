package main

import (
	"fmt"
)

//a unica estrutura de laço em go é o for

func main() {
	i := 1

	//similar ao laço while
	for i <= 10 {
		fmt.Print(i, ", ")
		i++
	}

	fmt.Println("")

	for i := 0; i <= 20; i += 2 {
		fmt.Printf("%d, ", i)
	}

	fmt.Println("")

	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Printf("Par: %d, ", i)
		} else {
			fmt.Printf("Impar: %d, ", i)
		}
	}

	/* laço infinito
	for {

	}
	*/
	fmt.Println("")

	for {
		fmt.Println("saindo do laço")
		break
	}
}
