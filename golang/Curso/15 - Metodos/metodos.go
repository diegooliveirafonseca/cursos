package main

import "fmt"

type portador struct {
	nome   string
	cpf    string
	cartao cartao
}

type cartao struct {
	numero         string
	limite         float32
	saldo          float32
	dataVencimento string
}

func (p *portador) criarCartao() {
	var cartao1 cartao
	//p3 := pessoa{nome: "Pérola", idade: 7}
	cartao1.numero = "5162928031941383"
	cartao1.dataVencimento = "0430"
	cartao1.limite = 1000.0
	p.cartao = cartao1
	fmt.Println(cartao1)
}

func (p *portador) adicionarSaldo(valor float32) {
	p.cartao.saldo += valor
	fmt.Printf("Saldo de %f adicionado no cartao %s\n", valor, p.cartao.numero)
}

func main() {
	var portador1 portador
	portador1.nome = "Diego Oliveira"
	portador1.cpf = "00578187205"
	portador1.criarCartao()
	fmt.Println(portador1)
	portador1.adicionarSaldo(50.0)
	fmt.Println(portador1.cartao.saldo)
}
