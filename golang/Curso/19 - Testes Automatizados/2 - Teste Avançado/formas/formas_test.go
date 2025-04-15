package formas

import (
	"math"
	"testing"
)

func TestArea(t *testing.T) {
	t.Run("Retângulo", func(t *testing.T) {
		ret := Retangulo{10, 12}
		areaEsperada := float64(120)
		areaRecebida := ret.Area()

		if areaRecebida != areaEsperada {
			t.Fatalf("A área recebida %f é diferente da área esperada %f", areaRecebida, areaEsperada)
		}
	})

	t.Run("Círculo", func(t *testing.T) {
		ret := Circulo{10}
		areaEsperada := float64(math.Pi * math.Pow(ret.raio, 2))
		areaRecebida := ret.Area()

		if areaRecebida != areaEsperada {
			t.Fatalf("A área recebida %f é diferente da área esperada %f", areaRecebida, areaEsperada)
		}
	})
}
