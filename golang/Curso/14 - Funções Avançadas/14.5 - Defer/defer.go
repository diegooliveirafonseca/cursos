package main

import "fmt"

// defer serve pare adiar a executação do comando no qual ele está prefixando até o ultimo momento possível.
// Se for em uma função sem retorno, será a ultima coisa a ser executada. Se tiver retorno, será executada antes do retorno.
func funcao1() {
	fmt.Println("Executando função 1")
}

func funcao2() {
	fmt.Println("Executando função 2")
}

func funcao3() {
	fmt.Println("Executando função 3")
}

func calcularMedia(n1, n2 float32) bool {
	defer fmt.Println("Média calculada, retornando resultado!")
	fmt.Println("Entrando na função de média!")
	media := (n1 + n2) / 2
	if media >= 6 {
		return true
	}
	return false
}

func main() {
	//funcao1()
	//defer funcao2()
	//funcao3()
	fmt.Println(calcularMedia(8, 7))

}
