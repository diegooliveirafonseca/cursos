package main

import "fmt"

func main() {
	fmt.Println("Loops")

	//Tipos de for

	//For com iteração interna
	i := 0
	for i < 10 {
		i++
		//fmt.Println(i)
	}

	//for com indice, condição e iteração no cabeçalho do for
	for j := 0; j < 10; j++ {
		//fmt.Println(j)
	}

	//for de arrays e slice
	//Caso não queira mostrar o indice, é só colocar um underline(_) no lugar dele

	//nomes := [3]string{"João", "Davi", "Lucas"}

	//for indice, nome := range nomes {
	//	fmt.Println(indice, nome)
	//}

	// marcasCarro := []string{"Ford", "Fiat", "Wolkswagen"}

	// for _, nome := range marcasCarro {
	// 	fmt.Println(nome)
	// }

	//for para maps

	usuario := map[string]string{
		"nome":      "Diego",
		"sobrenome": "Fonseca",
	}

	for chave, valor := range usuario {
		fmt.Println(chave, valor)
	}

	//não existe loops de structs
}
