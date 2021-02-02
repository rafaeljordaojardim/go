package enderecos

import "strings"

// TipoDeEndereco verifica se endereço tem um tipo valido e o retorna
func TipoDeEndereco(endereco string) string {
	tiposValidos := []string{"rua", "avenida", "estrada", "rodovia"}
	enderecoLetraMinuscula := strings.ToLower(endereco)
	primeiraPalavra := strings.Split(enderecoLetraMinuscula, " ")[0]

	enderecoTemUmTipoValido := false
	for _, tipo := range tiposValidos {
		if tipo == primeiraPalavra {
			enderecoTemUmTipoValido = true
		}
	}

	if enderecoTemUmTipoValido {
		return strings.Title(primeiraPalavra)
	}

	return "Tipo inválido"
}
