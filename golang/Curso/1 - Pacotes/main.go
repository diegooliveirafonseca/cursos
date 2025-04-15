package main

import (
	"fmt"
	"modulo/auxiliar"

	"github.com/badoux/checkmail"
	"github.com/klassmann/cpfcnpj"
)

func main() {
	fmt.Println("Escrevendo do arquivo main")
	auxiliar.Escrever()

	erro := checkmail.ValidateFormat("devbook@@gmail.com")
	fmt.Println(erro)

	erro2 := cpfcnpj.ValidateCPF("12345678910")
	fmt.Println(erro2)
}
