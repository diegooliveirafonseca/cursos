package main

import "fmt"

func calculosMatematicos(numero1, numero2 int) (soma, subtracao int, divisao, multiplicacao float32) {
	soma = numero1 + numero2
	subtracao = numero1 - numero2
	divisao = float32(numero1) / float32(numero2)
	multiplicacao = float32(numero1) * float32(numero2)
	return
}

func main() {

	soma, subtracao, divisao, multiplicacao := calculosMatematicos(20, 20)
	fmt.Println(soma, subtracao, divisao, multiplicacao)
}
