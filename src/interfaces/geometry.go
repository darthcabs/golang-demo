package main

import (
	"fmt"
	"math"
)

type forma interface {
	area() float64
	perimetro() float64
}

// Retângulo
type retangulo struct {
	largura, altura float64
}

func (r retangulo) area() float64 {
	return r.largura * r.altura
}

func (r retangulo) perimetro() float64 {
	return 2*r.largura + 2*r.altura
}

// Círculo
type circulo struct {
	raio float64
}

func (c circulo) area() float64 {
	return math.Pi * c.raio * c.raio
}

func (c circulo) perimetro() float64 {
	return 2 * math.Pi * c.raio
}

// O poder das interfaces!
func medir(f forma) {
	fmt.Printf("%T: %v\n", f, f)
	fmt.Printf("Área: %.1f\n", f.area())
	fmt.Printf("Perímetro: %.1f\n\n", f.perimetro())
}

func main() {
	r := retangulo{largura: 3, altura: 4}
	c := circulo{raio: 5}

	medir(r)
	medir(c)
}
