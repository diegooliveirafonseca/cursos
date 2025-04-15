package main

import (
	"fmt"
	"introducao-teste/enderecos"
)

func main() {
	tipoEndereco := enderecos.TipoDeEndereco("Rua dos Bobos")
	fmt.Println(tipoEndereco)
}
