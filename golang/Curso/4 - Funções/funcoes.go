package main

import "fmt"

func somar(n1 int8, n2 int8) int8 {
	return n1 + n2
}

func calculosMatematicos(n1, n2 int8) (int8, int8, float32, int16) {
	soma := n1 + n2
	subtracao := n1 - n2
	divisao := n1 / n2
	multiplicacao := n1 * n2
	return soma, subtracao, float32(divisao), int16(multiplicacao)
}

func musiquinha(nome string) (string, string, string, string, string) {
	return "Ha Ra Hu Ru ", nome, ", eu vou comer o seu ", "*", "bolo"

}

func main() {
	soma := somar(10, 20)
	fmt.Println(soma)

	var f = func(txt string) string {
		fmt.Println(txt)
		return "resultado: " + txt
	}

	resultado := f("hello world")
	fmt.Println(resultado)

	resultadoSoma, resultadoSub, resultadoDiv, resultadoMult := calculosMatematicos(20, 2)
	fmt.Println("Soma: ", resultadoSoma, " | Sub: ", resultadoSub, " | Div: ", resultadoDiv, " | Mult: ", resultadoMult)

	//retorno de funções podem ser ignorados colocando um underline(_) no lugar da variavel de retorno
	//Exemplo:
	parte1, parte2, parte3, _, parte5 := musiquinha("Rafaela")
	fmt.Println(parte1, parte2, parte3, parte5)

}
