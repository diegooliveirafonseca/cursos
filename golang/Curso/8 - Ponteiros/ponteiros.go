package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println("Ponteiros")

	//Criação de variavel com copia de valor
	var variavel1 int = 100
	var variavel2 int = variavel1
	fmt.Println(variavel1, variavel2)

	variavel1++
	fmt.Println(variavel1, variavel2)

	//O ponteiro não é uma copia de valor, é uma referenciação  de memoria
	//para criar um ponteiro usa a declaração normal, mas quando vai sertar o tipo, utiliza um asteristico (*) no tipo.

	var variavel3 int
	var ponteiro *int

	fmt.Println(variavel3, ponteiro, reflect.TypeOf(variavel3), reflect.TypeOf(ponteiro))

	//Para fazer a copia do endereço de memoria, precisa usar o  & antes da variavel
	variavel3 = 10
	ponteiro = &variavel3

	fmt.Println(variavel3, ponteiro, reflect.TypeOf(variavel3), reflect.TypeOf(ponteiro))

	//Para ver o valor que está naquele endereço de memoria, precisa usar o * antes do nome da variavel de ponteiro -> Isso se chama DESREFERENCIAÇÃO
	fmt.Println(variavel3, *ponteiro, reflect.TypeOf(variavel3), reflect.TypeOf(ponteiro))
	variavel3 = 15
	fmt.Println(variavel3, *ponteiro, reflect.TypeOf(variavel3), reflect.TypeOf(ponteiro))
}
