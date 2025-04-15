package main

import (
	"fmt"
	"time"
)

func escrever(texto string, canal chan string) {
	for i := 0; i < 5; i++ {
		fmt.Println(i)
		canal <- texto
		time.Sleep(time.Second)
	}
	close(canal)
}

func main() {
	canal := make(chan string)

	go escrever("Olá mundo", canal)

	/*for { Foi substituido pelo range do canal. Enquanto tiver itens no canal, a execução continua, no momento que não tiver mais, o programa pula para a proxima execução
		mensagem, aberto := <-canal
		fmt.Println(aberto)
		if !aberto {
			break
		}
		fmt.Println(mensagem)
	}*/
	for mensagem := range canal {
		fmt.Println(mensagem)
	}

	fmt.Println("Fim do programa ")

}
