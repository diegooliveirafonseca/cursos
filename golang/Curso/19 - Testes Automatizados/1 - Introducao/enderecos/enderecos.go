package enderecos

import (
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func TipoDeEndereco(endereco string) string {
	tiposValidos := []string{"rua", "avenida", "rodovia", "estrada"}

	enderecoEmLetraMinuscula := strings.ToLower(endereco)
	primeiraPalavraDoEndereco := strings.Split(enderecoEmLetraMinuscula, " ")[0]

	enderecoTemUmTipoValido := false

	for _, tipo := range tiposValidos {
		if primeiraPalavraDoEndereco == tipo {
			enderecoTemUmTipoValido = true
		}
	}

	if enderecoTemUmTipoValido {
		return cases.Title(language.Und, cases.NoLower).String(primeiraPalavraDoEndereco)
	}
	return "Tipo Inválido!"
}
