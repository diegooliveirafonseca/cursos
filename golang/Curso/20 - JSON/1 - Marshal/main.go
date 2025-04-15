package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
)

type cachorro struct {
	Nome string `json:"nome"`
	Raca string `json:"raca"`
}

func main() {
	c1 := cachorro{"Melky", "Dog Alemão"}
	fmt.Println(c1)

	cachorroEmJSON, erro := json.Marshal(c1)
	if erro != nil {
		log.Fatal(erro)
	}
	fmt.Println(cachorroEmJSON)
	fmt.Println(bytes.NewBuffer(cachorroEmJSON))

	c2 := map[string]string{
		"nome": "Nika",
		"raca": "Dog",
	}

	cachorro2EmJSON, erro := json.Marshal(c2)
	fmt.Println(cachorro2EmJSON)

	fmt.Println(bytes.NewBuffer(cachorro2EmJSON))
}
