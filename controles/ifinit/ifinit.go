package main

import (
	"fmt"
	"math/rand"
	"time"
)

func numeroAleatorio() int {
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)
	return r.Intn(10)
}

func main() {

	//inicialização de variavel visivel apenas no bloco if else
	if i := numeroAleatorio(); i > 5 {
		fmt.Println("Ganhou com o número", i)
	} else {
		fmt.Println("Perdeu com o número", i)
	}
}
