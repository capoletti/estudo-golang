package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int) //canal sem buffer

	go rotina(c)

	fmt.Println("antes da leitura")
	fmt.Println(<-c)
	fmt.Println("fim")
	fmt.Println(<-c) //gera o deadlock

}

func rotina(c chan int) {
	time.Sleep(time.Second)
	c <- 1 //operação bloqueante
	fmt.Println("depois da leitura")
}
