package main

import (
	"fmt"

	"github.com/capoletti/html"
)

func main() {

	c := unirMensagens(
		html.Titulo("https://www.google.com", "https://www.youtube.com"),
		html.Titulo("http://www.uol.com", "https://www.amazon.com"),
	)

	fmt.Println(<-c)
	fmt.Println(<-c)
	fmt.Println(<-c)
	fmt.Println(<-c)

}

//canal somente leitura <-chan
func encaminhar(origem <-chan string, destino chan string) {
	for {
		destino <- <-origem
	}
}

//multiplexar, unir mensagens de multiplos canais no mesmo
func unirMensagens(entrada1, entrada2 <-chan string) <-chan string {
	c := make(chan string)

	go encaminhar(entrada1, c)
	go encaminhar(entrada2, c)
	return c
}
