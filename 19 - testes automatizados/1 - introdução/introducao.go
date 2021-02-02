package main

import (
	"fmt"
	"introducao-testes/enderecos"
)

func main() {
	tipoDeEndereco := enderecos.TipoDeEndereco("Paulista")
	fmt.Println(tipoDeEndereco)
}
