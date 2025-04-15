package main

import (
	"errors"
	"fmt"
)

func main() {
	//Inteiros
	//int8, int16, int32, int64
	//int -> utiliza a arquitetura do computador, se o pc for 32 bits ou 64bis, ele utiliza isso para o numero int
	//uint //unsygned int -> Só funciona com numeros sem sinal -1000 da erro, 1000 funciona. segue a mesma conversão do int -> uint8 uint16....
	var numero int8 = 100
	numero2 := 100000
	var numero3 uint = 10000

	//alias para int
	//int32 -> rune
	//uint8 -> byte

	var numero4 byte = 123

	//float32, float34, não pode usar float para declarar, se for usar inferencia(:= sem tipo) ele utiliza a arquitetura do computador
	var numeroReal1 float32 = 123.45
	var numeroReal2 float64 = 12300000000.45
	numeroReal3 := 5455688455.45

	//valor inicial de variavies
	// string -> vazio
	//int e  float -> 0
	var valorFloat float64
	fmt.Println(valorFloat)
	// bool -> false
	//error -> nil

	var x float32

	//Strings -> normal. para criar um
	var str string = "abc"

	//booleano -> true, false
	var ok bool = true

	//error
	var erro error
	var erroInterno error = errors.New("Erro Interno")

	fmt.Println(numero, numero2, numero3, numero4, numeroReal1, numeroReal2, numeroReal3, x, str, ok, erro, erroInterno)
}
