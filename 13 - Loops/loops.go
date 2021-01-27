package main

import (
	"fmt"
)

func main() {
	i := 0

	for i < 10 {
		i++
		// fmt.Println("Incrementando i")
		// time.Sleep(time.Second)
	}

	for j := 0; j < 10; j++ {
		// fmt.Println("Incrementando j ", j)
		// time.Sleep(time.Second)
	}

	nomes := [3]string{"João", "Davi", "Pedro"}

	for indice, nome := range nomes {
		fmt.Println(indice, nome)
	}

	usuario := map[string]string{
		"nome":      "Rafael",
		"sobrenome": "Jardim",
	}

	for chave, valor := range usuario {
		fmt.Println(chave, valor)
	}
}
