package main

import "fmt"

func main() {
	fmt.Println("Estruturas de Controle")

	numero := -9
	if numero > 0 {
		fmt.Println("Maior que 0")
	} else if numero < -10 {
		fmt.Println("Menor que -10")
	} else {
		fmt.Println("Maior que -10 e menor que 0")
	}

	nome := 30

	if nome < 20 {
		return
	}
	nome++
	fmt.Println(nome)
	//If init

	if outroNumero := numero; outroNumero == -9 {
		fmt.Println("Igual a -9")
	}
}
