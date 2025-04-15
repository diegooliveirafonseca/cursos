package main

import "fmt"

func main() {
	var nome string = "Diego"
	sobrenome := "Oliveira"
	var (
		cpf string = "12345678910"
		rg  string = "6543210"
	)
	escreveCom, chutaCom := "Direita", "Esquerda"
	const altura string = "1,78"
	fmt.Println(nome, sobrenome, cpf, rg, escreveCom, chutaCom, altura)
	escreveCom, chutaCom = chutaCom, escreveCom
	fmt.Println(escreveCom, chutaCom)
}
