package main

import (
	"fmt"
	"time"
)

func fale(pessoa, texto string, quantidade int) {

	for i := 0; i < quantidade; i++ {
		time.Sleep(time.Second)
		fmt.Printf("%s %s (iteração %d) \n", pessoa, texto, i+1)
	}

}

func main() {
	//fale("Maria", "Pq vc não fala comigo?", 3)
	//fale("João", "Só posso falar depois de você", 1)

	go fale("Maria", "Ei...", 5)
	go fale("João", "Opa...", 5)

	time.Sleep(time.Second * 5)
}
