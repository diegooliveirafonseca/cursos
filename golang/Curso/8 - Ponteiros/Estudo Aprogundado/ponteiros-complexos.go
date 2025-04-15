package main

import "fmt"

func main() {
	x := 10
	p := &x // Ponteiro recebe o endereço de x

	fmt.Println("Endereço de x:", p) // Exibe o endereço de x
	fmt.Println("Valor de x:", *p)   // Exibe o valor de x (desreferenciação)

	*p = 20                            // Modifica o valor de x através do ponteiro
	fmt.Println("Novo valor de x:", x) // Agora x vale 20
}
