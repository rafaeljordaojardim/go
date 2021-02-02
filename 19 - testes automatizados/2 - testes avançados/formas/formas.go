package formas

import (
	"math"
)

// Retangulo faz algo
type Retangulo struct {
	Altura  float64
	Largura float64
}

// Area faz algo
func (r Retangulo) Area() float64 {
	return r.Altura * r.Largura
}

// Area faz algo
func (c Circulo) Area() float64 {
	return math.Pi * math.Pow(c.Raio, 2)
}

// Circulo faz algo
type Circulo struct {
	Raio float64
}

// Forma faz algo
type Forma interface {
	area() float64
}
