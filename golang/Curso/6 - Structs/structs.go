package main

import "fmt"

type pessoa struct {
	nome     string
	cpf      string
	rg       string
	idade    uint
	endereco endereco
}

type endereco struct {
	logradouro  string
	numero      string
	complemento string
}

func main() {
	fmt.Println("Arquivo structs")

	var p1 pessoa
	p1.nome = "Diego Oliveira"
	p1.cpf = "00578187205"
	p1.rg = "6298033"
	p1.idade = 34
	fmt.Println(p1)

	end1 := endereco{"Rua Pampulha", "10", "Kit net 5"}

	p2 := pessoa{"Rafaela", "01874718270", "123456", 33, end1}
	fmt.Println(p2)

	p3 := pessoa{nome: "Pérola", idade: 7}
	fmt.Println(p3)

}
