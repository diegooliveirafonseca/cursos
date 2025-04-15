package main

import "fmt"

type participante struct {
	codigo string
}
type pessoa struct {
	participante
	cpf string
	rg  string
}
type organizacao struct {
	participante
	cnpj              string
	inscricaoEstadual string
}

func main() {
	fmt.Println("Herança")

	p1 := participante{codigo: "1"}
	fmt.Println(p1)
	pessoa1 := pessoa{participante: p1, cpf: "00578187205"}
	fmt.Println(pessoa1)
	fmt.Println(pessoa1.codigo)

}
