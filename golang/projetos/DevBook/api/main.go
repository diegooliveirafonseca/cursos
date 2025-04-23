package main

import (
	"api/src/config"
	"api/src/router"
	"fmt"
	"log"
	"net/http"
)

//Chave ja foi gerada - Comentando a função init
/*func init() {
	//Cria um slice de bytes de 64 bits
	chave := make([]byte, 64)

	//Preenche a chave com valores aleatorios
	if _, erro := rand.Read(chave); erro != nil {
		log.Fatal(erro)
	}
	//Converte o slice de bytes de numeros aleatorios em string
	stringBase64 := base64.StdEncoding.EncodeToString(chave)
	fmt.Println(stringBase64)
}*/

func main() {
	config.Carregar()
	r := router.Gerar()

	fmt.Printf("Escutando na porta %d\n", config.Porta)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Porta), r))
}
