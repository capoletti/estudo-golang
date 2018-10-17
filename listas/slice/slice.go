package main

import (
	"fmt"
	"reflect"
)

func main() {

	a1 := [3]int{1, 2, 3} //array
	s1 := []int{1, 2, 3}  //slice

	fmt.Println(reflect.TypeOf(a1), a1)
	fmt.Println(reflect.TypeOf(s1), s1)

	a2 := [5]int{1, 2, 3, 4, 5}
	s2 := a2[1:3]
	a2[1] = 7

	fmt.Println(a2)
	fmt.Println(s2)

	//slice possui um tamanho e um ponteiro para um elemento de um array
	//criando uma estrutura continua do elemento do ponteiro até o final do tamanho
	s3 := a2[:2]
	s4 := s3[1:2]
	fmt.Println(s3, s4)

}
