package main

import "fmt"

func diaDaSemana(numero int) string {
	switch numero {
	case 1:
		return "Domingo"
	case 2:
		return "Segunda-Feira"
	case 3:
		return "Terça-Feira"
	case 4:
		return "Quarta-Feira"
	case 5:
		return "Quinta-Feira"
	case 6:
		return "Sexta-Feira"
	case 7:
		return "Sabado"
	default:
		return "Número Invalido"
	}

}

func sexo(sexo string) string {
	var retorno string
	switch {
	case sexo == "M":
		retorno = "Masculino"
	case sexo == "F":
		retorno = "Feminino"
	default:
		retorno = "Sexo Inválido"
	}
	return retorno
}

func main() {
	fmt.Println("Switch")
	dia := diaDaSemana(1)
	fmt.Println(dia)

	sexo := sexo("M")
	fmt.Println(sexo)

}
