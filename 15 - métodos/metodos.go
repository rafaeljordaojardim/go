package main

import "fmt"

type usuario struct {
	nome  string
	idade uint8
}

func (u usuario) salvar() {
	fmt.Println("Dentro do méodo salvar")
	fmt.Println(u.nome, u.idade)
}

func (u *usuario) fazerAniversario() {
	u.idade++
}

func main() {
	usuario1 := usuario{"usuario", 20}

	fmt.Println(usuario1)
	usuario1.salvar()
}
