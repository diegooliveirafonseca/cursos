package main

import "fmt"

func main() {
	fmt.Println("Maps")

	usuario := map[string]string{
		"nome": "Diego",
	}
	fmt.Println(usuario)

	//Maps aninhados

	usuario2 := map[string]map[string]string{
		"nome": {
			"primeiro": "Diego",
			"ultimo":   "Fonseca",
		},
		"data-nascimento": {
			"dia": "04",
			"mes": "Abril",
			"ano": "1990",
		},
		"emprego": {
			"empresa": "Acception",
			"cargo":   "Analista de Sistemas",
		},
	}
	fmt.Println(usuario2)

	//Deletar um item do map

	delete(usuario2, "emprego")
	fmt.Println(usuario2)

	//Adicionar um item no map

	usuario2["signo"] = map[string]string{
		"nome": "Aries",
	}
	fmt.Println(usuario2)
}
