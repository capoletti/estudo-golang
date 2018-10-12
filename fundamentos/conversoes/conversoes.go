package main

import (
	"fmt"
	"strconv"
)

func main() {

	x := 2.4
	y := 2

	fmt.Println(x / float64(y))

	nota := 6.9
	notaFinal := int(nota)

	fmt.Println(notaFinal)

	var a = 'a'

	fmt.Println("Teste: " + string(123))
	fmt.Println("Teste: " + string(97))
	fmt.Println("Teste: " + string(a))

	fmt.Println("Inteiro para string: " + strconv.Itoa(123))
	fmt.Println("Char para string: " + strconv.QuoteRune(a))

	num, err := strconv.Atoi("123a")

	fmt.Println(num)
	fmt.Println(err)

	b, _ := strconv.ParseBool("true")
	fmt.Println(b)

	b1, _ := strconv.ParseBool("1")
	fmt.Println(b1)

}
