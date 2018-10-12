package main

import (
	"fmt"
	"math"
	"reflect"
)

func main() {

	fmt.Println(1, 2, 1000)
	fmt.Println("Literal inteiro é: ", reflect.TypeOf(32000))

	//uint8 .. 64 são apenas positivo (sem sinal)
	var b byte = 255
	fmt.Println("o byte é: ", reflect.TypeOf(b))

	i1 := math.MaxInt64
	fmt.Println("maximo int64: ", i1)

	var i2 rune = 'a'
	fmt.Println("o tipo i2 é: ", reflect.TypeOf(i2))
	fmt.Println("valor do i2: ", i2)

	//string é tipo primitivo
	s1 := "variavel do tipo string"
	fmt.Println(s1)
	fmt.Println(len(s1))

	s2 := `variavel
	do
	tipo
	string`

	fmt.Println(s2)
	fmt.Println(len(s2))

	//em go não há o tipo char
	char := 'a'
	fmt.Println("o tipo char é: ", reflect.TypeOf(char))
	fmt.Println("valor do char: ", char)
}
