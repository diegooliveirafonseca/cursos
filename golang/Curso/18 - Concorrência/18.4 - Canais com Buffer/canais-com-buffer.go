package main

import "fmt"

func main() {
	canal := make(chan string, 3) //Buffer= define uma capacidade para o canal, desta forma evita o travamento do canal no envio  ou no recebimento por sabermos a sua capacidade

	canal <- "Olá Mundo!"
	canal <- "Programando em GO!"
	canal <- "Mensagem 3!"

	mensagem := <-canal
	fmt.Println(mensagem)
	mensagem2 := <-canal
	fmt.Println(mensagem2)
	mensagem3 := <-canal
	fmt.Println(mensagem3)
}
