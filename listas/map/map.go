package main

import (
	"fmt"
)

func main() {

	map1()
	map2()
	map3()
}

func map1() {

	aprovados := make(map[int]string)

	aprovados[12345678977] = "Maria"
	aprovados[12345678978] = "Pedro"
	aprovados[12345678979] = "Ana"

	fmt.Println(aprovados)

	for cpf, nome := range aprovados {
		fmt.Printf("%s (CPF: %d) \n", nome, cpf)
	}

	fmt.Println(aprovados[12345678979])
	delete(aprovados, 12345678977)
	fmt.Println(aprovados)
}

func map2() {

	funcESalario := map[string]float64{
		"João":  1200.00,
		"Maria": 2500.20,
		"Jose":  2300.00,
	}

	funcESalario["Zezinho"] = 5000.40
	delete(funcESalario, "Não existe") //ao excluir elemento inexistente não há ocorre erro

	fmt.Println(funcESalario)

}

func map3() {

	funcsPorLetra := map[string]map[string]float64{
		"J": {
			"João": 1200.00,
			"Jose": 1250.00,
		},
		"G": {
			"Gabriel": 1300.00,
			"Gustavo": 1240.00,
		},
		"P": {
			"Pedro": 1210.00,
		},
	}

	fmt.Println(funcsPorLetra)

	for letra, funcs := range funcsPorLetra {
		for nome, salario := range funcs {
			fmt.Printf("%s: %s (%.2f) \n", letra, nome, salario)
		}
	}

}
