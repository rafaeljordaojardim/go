package main

import "fmt"

type endereco struct {
	logradouro string
	numero     uint8
}

type usuario struct {
	nome     string
	idade    uint8
	endereco endereco
}

func main() {
	fmt.Println("Arquivo structs")

	var u usuario
	u.nome = "Rafael"
	u.idade = 21

	u2 := usuario{nome: "Rafael", idade: 21}
	fmt.Println(u2)

	u3 := usuario{nome: "Rafael"}
	fmt.Println(u3)

}
