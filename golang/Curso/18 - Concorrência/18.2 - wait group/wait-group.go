package main

import (
	"fmt"
	"sync"
	"time"
)

func escrever(texto string) {
	for i := 0; i < 5; i++ {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}

func main() {

	var waitGroup sync.WaitGroup

	waitGroup.Add(3) //Diz que tem duas goroutines para rodar

	go func() {
		go escrever("Olá mundo")
		waitGroup.Done() // waitGroup -1
	}()

	go func() {
		go escrever("Programando em Go")
		waitGroup.Done() // waitGroup -1
	}()

	go func() {
		escrever("Escrevendo")
		waitGroup.Done() // waitGroup -1
	}()
	waitGroup.Wait() //Vai aguardar as duas goroutines terminarem para finalizar o programa.
}
