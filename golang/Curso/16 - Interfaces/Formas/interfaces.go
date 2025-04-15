package main

import (
	"fmt"
	"math"
)

type forma interface {
	area() float64
}

type retangulo struct {
	altura  float64
	largura float64
}

func (r retangulo) area() float64 {
	return r.altura * r.largura
}

type circulo struct {
	raio float64
}

// pacote math tem varias funções matemáticas
// math.Pi trás o valor de PI
// math.Pow faz o calculo de numero com potencia, sendo o primeiro parametro o numero e o segundo a potencia
func (c circulo) area() float64 {
	return math.Pi * math.Pow(c.raio, 2)
}

func escreverArea(f forma) {
	fmt.Printf("O valor da área é: %f\n", f.area())
}

func main() {
	r := retangulo{10, 15}
	escreverArea(r)
	c := circulo{10}
	escreverArea(c)
}
