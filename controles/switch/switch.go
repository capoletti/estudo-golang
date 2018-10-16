package main

import (
	"fmt"
	"time"
)

func notaConceito(n float64) string {
	var nota = int(n)

	switch nota {
	case 10:
		fallthrough //executar o próximo case
	case 9:
		return "A"
	case 8, 7:
		return "B"
	case 6, 5:
		return "C"
	case 4, 3:
		return "D"
	case 2, 1, 0:
		return "E"
	default:
		return "inválida"
	}
}

func notas() {

	for i := 11.0; i >= -1; i-- {
		fmt.Println(i, notaConceito(i))
	}
}

func tempo() {

	t := time.Now()

	//procura o primeiro case verdadeiro
	switch {
	case t.Hour() < 12:
		fmt.Println("Bom dia!")
	case t.Hour() < 18:
		fmt.Println("Boa tarde!")
	default:
		fmt.Println("Boa noite!")

	}
}

func tipo(i interface{}) string {
	switch i.(type) {

	case int:
		return "inteiro"
	case float32, float64:
		return "real"
	case string:
		return "string"
	case func():
		return "função"
	default:
		return "desconhecido"
	}
}

func tipos() {
	fmt.Println(tipo(2.3))
	fmt.Println(tipo(2))
	fmt.Println(tipo("opa"))
	fmt.Println(tipo(func() {}))
	fmt.Println(tipo(time.Now()))
}

func main() {

	notas()
	tempo()
	tipos()

}
