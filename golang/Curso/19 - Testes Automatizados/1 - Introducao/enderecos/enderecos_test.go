// TESTE DE UNIDADE
package enderecos_test

import (
	. "introducao-teste/enderecos"
	"testing"
)

type cenarioDeTeste struct {
	enderecoInserido string
	retornoEsperado  string
}

func TestTipoDeEndereco(t *testing.T) {

	t.Parallel()

	enderecoParaTeste := "Rua dos Bobos"
	enderecoEsperado := "Rua"
	enderecoRecebido := TipoDeEndereco(enderecoParaTeste)

	if enderecoRecebido != enderecoEsperado {
		t.Errorf("O tipo recebido é diferente do esperado! Esperado: %s recebido %s", enderecoEsperado, enderecoRecebido)
	}
}

func TestTipoDeEndereco2(t *testing.T) {

	t.Parallel()

	cenariosDeTeste := []cenarioDeTeste{
		{"Rua Borborema", "Rua"},
		{"Avenida Centenário", "Avenida"},
		{"Estrada da Maracacuera", "Estrada"},
		{"Rodovia Augusto Montenegro", "Rodovia"},
		//{"Praça da República", "Tipo Inválido!"},
		//{"", "Tipo Inválido!"},
	}

	for _, cenario := range cenariosDeTeste {
		retornoRecebido := TipoDeEndereco(cenario.enderecoInserido)
		if retornoRecebido != cenario.retornoEsperado {
			t.Errorf("O tipo recebido %s é diferente do esperado %s\n", retornoRecebido, cenario.retornoEsperado)
		}

	}
}
