package enderecos

import "testing"

type cenarioDeTeste struct {
	enderecoInserido string
	retornoEsperado  string
}

func TestTipoDeEndereco(t *testing.T) {
	cenariosDeTeste := []cenarioDeTeste{
		{"Rua ABC", "Rua"},
		{"Avenida Paulista", "Avenida"},
		{"Rodovia dos Imigrantes", "Rodovia"},
		{"Praça das Rosas", "Tipo Inválido"},
		{"Estrada ABC", "Estrada"},
		{"ESTRADA ABC", "Estrada"},
		{"RUA ABC", "Rua"},
		{"AVENIDA Paulista", "Avenida"},
		{"RODOVIA dos Imigrantes", "Rodovia"},
		{"", "Tipo Inválido"},
	}

	for _, cenario := range cenariosDeTeste {
		tipoDeEnderecoRecebido := TipoEnderecos(cenario.enderecoInserido)

		if tipoDeEnderecoRecebido != cenario.retornoEsperado {
			t.Errorf("O tipo de endereço recebido foi diferente do esperado. Esperado: %s, Recebido: %s", cenario.retornoEsperado, tipoDeEnderecoRecebido)
		}
	}
}
