package main

import "fmt"

func soma(numeros ...int) int {
	total := 0
	for _, numero := range numeros {
		total += numero
	}
	return total
}

func escrever(x int, texto string, numeros ...int) {
	for _, numero := range numeros {
		fmt.Println(x, texto, numero)
	}
}

func main() {
	totalDaSoma := soma(1, 2, 3, 4, 5, 6)
	fmt.Println(totalDaSoma)

	escrever(4, "Olá Mundo", 1, 3, 5, 7, 9)
}
