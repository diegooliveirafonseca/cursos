package main

import "fmt"

// A função panic serve para parar a execução do programa caso sua condição seja aceita. Para evitar de o programa parar, se usa a função recover
func recuperarExecucao() {
	if r := recover(); r != nil {
		fmt.Println("Execução recuperada")
	}
}

func alunoEstaAprovado(n1, n2 float64) bool {
	defer recuperarExecucao()
	media := (n1 + n2) / 2
	if media > 6 {
		return true
	} else if media < 6 {
		return false
	}
	panic("A MÉDIA É EXATAMENTE 6!")
}

func main() {
	fmt.Println(alunoEstaAprovado(6, 6))
	fmt.Println("Pós Execução")
}
