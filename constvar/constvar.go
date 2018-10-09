package main

import (
	"fmt"
	"math"
)

func main() {
	const PI float64 = 3.1415
	var raio = 3.2 //tipo float64

	//forma reduzida de
	area := PI * math.Pow(raio, 2)

	fmt.Println("area:", area)

	const (
		a = 1
		b = 2
	)

	var (
		c = 3
		d = 4
	)

	fmt.Println(a, b, c, d)

	var e, f = true, false

	fmt.Println(e, f)

	g, h, i := 2, false, "valor"

	fmt.Println(g, h, i)

}
