package main

import "fmt"

func fibonacci(posicao uint) uint {
	//condição de parada
	fmt.Println("posicao", posicao)
	if posicao <= 1 {
		return posicao
	}
	retorno := fibonacci(posicao-2) + fibonacci(posicao-1)
	fmt.Println("retorno:", retorno)
	return retorno
}

func main() {

	//1,1,2,3,5,8,13,21,32,55

	posicao := uint(3)
	fmt.Println(fibonacci(posicao))
	fmt.Println("-------------------")
	//for i := uint(1); i <= posicao; i++ {
	//	fmt.Println(fibonacci(i))
	//}

}
