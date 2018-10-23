package main

import "math"

//inicializando com letra maiúscula é Público
//inicializando com letra miniscula é Privado
//visibilidade privada é no pacote ao invés das estruturas

//Ponto representa uma coordenada
type Ponto struct {
	x float64
	y float64
}

func catetos(a, b Ponto) (cx, cy float64) {
	cx = math.Abs(b.x - a.x)
	cy = math.Abs(b.y - a.x)
	return
}

func Distancia(a, b Ponto) float64 {
	cx, cy := catetos(a, b)
	return math.Sqrt(math.Pow(cx, 2) + math.Pow(cy, 2))
}
