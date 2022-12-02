package main

import "math"

//Iniciando com letra maiúscula: Público (visível dentro e fora do pacote)
//Iniciando com letra minúscula: Privado (visível apenas dentro do pacote)

//Por exemplo...
//Ponto - gerará algo público
//ponto ou _Ponto - gerará algo privado

// Ponto representa uma coordenada no plano cartesiano
type Ponto struct {
	x float64
	y float64
}

func catetos(a, b Ponto) (cx, cy float64) {
	cx = math.Abs(a.x - b.x)
	cy = math.Abs(a.y - b.y)
	return
}

// Distancia é responsável por calcular a distância linear entre dois pontos.
func Distance(x, y Ponto) float64 {
	cx, cy := catetos(x, y)
	return math.Sqrt(math.Pow(cx, 2) + math.Pow(cy, 2))
}
