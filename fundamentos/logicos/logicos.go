package main

import (
	"fmt"
)

func comparar(item1, item2 bool) (bool, bool, bool) {
	ret1 := item1 && item2
	ret2 := item1 != item2 // ou exclusivo
	ret3 := item1 || item2

	return ret1, ret2, ret3
}

func main() {
	fmt.Println(comparar(true, true))
	fmt.Println(comparar(true, false))
	fmt.Println(comparar(false, true))
	fmt.Println(comparar(false, false))
}
