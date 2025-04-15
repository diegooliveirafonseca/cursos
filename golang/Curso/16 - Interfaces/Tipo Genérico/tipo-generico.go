package main

import "fmt"

func generica(interf interface{}) {
	fmt.Println(interf)
}

func main() {
	generica(1)
	generica("String")
	generica(true)
}

//A interface generica serve para burlar os tipos do go. mas só deve ser usado em casos especificos.
// O fmt.Println é uma interface generica.
