package main

import "fmt"

func inverteSinal(numero int) int {
	return numero * -1
}

func inverteSinalComPonteiro(numero *int) {
	*numero = *numero * -1
}

func main() {
	numero := 20
	fmt.Println("Valor Original: ", numero, "End. Memória: ", &numero)
	numeroRetorno := inverteSinal(numero)
	fmt.Println("Valor Invertido: ", numeroRetorno, "End. Memória: ", &numeroRetorno)

	numeroNovo := 40
	fmt.Println("Valor Original do ponteiro: ", numeroNovo, "End. Memória: ", &numeroNovo)
	inverteSinalComPonteiro(&numeroNovo)
	fmt.Println("Valor do ponteiro após inversão: ", numeroNovo, "End. Memória: ", &numeroNovo)
}
