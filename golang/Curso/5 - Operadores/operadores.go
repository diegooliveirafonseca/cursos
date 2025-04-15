package main

import "fmt"

func main() {
	//ARITMETICOS
	//+ - / * %
	soma := 2 + 1
	sub := 2 - 1
	div := 2 / 1
	mult := 2 * 1
	mod := 2 % 1

	fmt.Println(soma, sub, div, mult, mod)

	//ATRIBUICAO
	// = :=
	var v1 string = "v1"
	v2 := "v2"
	fmt.Println(v1, v2)

	//RELACIONAIS
	//> < >= <= == != retorno boleano

	//LOGICOS
	//&& || !
	var v3 bool
	fmt.Println(!v3)

	//UNARIOS
	//+= -= ++ -- *= /= %=

	//TERNARIO
	//Não existe no go

}
