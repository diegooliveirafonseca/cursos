package main

import "fmt"

//A função main é sempre executada primeiro que a função main.  normalmente usada para inicializar algum paramemtro ou configurar algo antes da execução do main

var n int

func init() {
	fmt.Println("Executando função init")
	n = 10
}

func main() {
	fmt.Println("Executando função main")
	fmt.Println(n)
}
