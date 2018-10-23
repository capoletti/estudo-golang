package main

import "fmt"

type imprimivel interface {
	toString() string
}

type pessoa struct {
	nome      string
	sobrenome string
}

type produto struct {
	nome  string
	preco float64
}

//interfaces sao implementadas implicitamente
func (p pessoa) toString() string {
	return p.nome + " " + p.sobrenome
}

func (p produto) toString() string {
	return fmt.Sprintf("%s: R$ %.2f", p.nome, p.preco)
}

func imprimir(x imprimivel) {
	fmt.Println(x.toString())
}

func main() {

	var item imprimivel

	item = pessoa{nome: "Jose", sobrenome: "Silva"}
	imprimir(item)

	item = produto{nome: "Carro", preco: 45679.98}
	imprimir(item)

	carro1 := ferrari{"F40", false, 0}
	fmt.Println(carro1)
	carro1.ligarTurbo()
	fmt.Println(carro1)

	var carro2 esportivo = &ferrari{"F45", false, 0} //necessita do &
	fmt.Println(carro2)
	carro2.ligarTurbo()
	fmt.Println(carro2)
}

//#02

type esportivo interface {
	ligarTurbo()
}

type ferrari struct {
	modelo          string
	turboLigado     bool
	velocidadeAtual int
}

func (f *ferrari) ligarTurbo() {
	f.turboLigado = true
}
