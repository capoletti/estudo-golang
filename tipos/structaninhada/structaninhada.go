package main

import (
	"fmt"
)

type item struct {
	produtoID  int
	quantidade int
	preco      float64
}

type pedido struct {
	userID int
	itens  []item
}

func (p pedido) valorTotal() float64 {
	total := 0.0
	for _, item := range p.itens {
		total += float64(item.quantidade) * item.preco
	}
	return total
}

func main() {
	pedido := pedido{
		userID: 1,
		itens: []item{
			item{produtoID: 1, quantidade: 3, preco: 5.20},
			item{produtoID: 1, quantidade: 5, preco: 7.43},
			item{produtoID: 1, quantidade: 10, preco: 12.80},
		},
	}

	fmt.Printf("Valor total R$ %.2f", pedido.valorTotal())
}
