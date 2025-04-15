package main

import "fmt"

func main() {
	fmt.Println("Arrays e Slices")

	//Arrays

	var array1 [5]string
	array1[0] = "Posicao 1"
	fmt.Println(array1)

	array2 := [5]string{"Posicao 1", "Posicao 2", "Posicao 3", "Posicao 4", "Posicao 5"}
	fmt.Println(array2)

	array3 := [...]int{1, 2, 3, 4, 5}
	fmt.Println(array3)

	//Slices
	slice := []int{10, 11, 12, 13, 14, 15, 16, 17}
	fmt.Println(slice)

	slice = append(slice, 18) //adicionar item ao slice
	fmt.Println(slice)

	slice2 := array2[1:3]
	fmt.Println(slice2)

	//ARRAY INTERNO
	// a função make precisa de 3 parametros: Tipo do Slice, Tamanho do slice(qtd de itens), e a capacidade(ate que tamanho ele pode ir)
	slice3 := make([]float32, 10, 11)
	fmt.Println(slice3)
	slice3 = append(slice3, 5)

	fmt.Println(slice3)
	fmt.Println(len(slice3)) //função len serve para ver o tamnho atual do slice
	fmt.Println(cap(slice3)) //função cap serve para ver a capacidade maxima  do slice.
	//Quando a capacidade maxima do slice é alcançada, o go cria um novo slice com a capacidade em dobro do slice anteriro para que o slice nunca tenha uma capacidade maxima

	slice4 := make([]float32, 5)
	fmt.Println(slice4)
	fmt.Println(len(slice4))
	fmt.Println(cap(slice4))

	slice4 = append(slice4, 10)
	fmt.Println(slice4)
	fmt.Println(len(slice4))
	fmt.Println(cap(slice4))
}
