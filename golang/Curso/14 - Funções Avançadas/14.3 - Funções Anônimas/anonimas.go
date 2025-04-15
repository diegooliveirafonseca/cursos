package main

import "fmt"

func main() {
	func() {
		fmt.Println("Olá  Mundo")
	}()

	func(texto string) {
		fmt.Println(texto)
	}("Passando Parâmetro")

	retorno := func(valor int, nome string) string {
		return fmt.Sprintf("Recebido -> %d de %s", valor, nome)
	}(10, "Diego")
	fmt.Println(retorno)
}
