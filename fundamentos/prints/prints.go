package main

import (
	"fmt"
)

func main() {

	fmt.Print("mesma")
	fmt.Print(" linha. ")

	fmt.Println("nova")
	fmt.Println(" linha. ")

	x := 3.141516

	//fmt.Println("valor: " + x)
	xs := fmt.Sprint(x)
	fmt.Println("valor: " + xs)
	fmt.Println("valor: ", x)

	fmt.Printf("valor formatada: %.2f", x)

	a := 1
	b := 1.9999
	c := false
	d := "valores"

	fmt.Printf("\n%d %f %.1f %t %s", a, b, b, c, d)
	fmt.Printf("\n%v %v %v %v", a, b, c, d)

}
