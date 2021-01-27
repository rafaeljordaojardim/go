package main

import "fmt"

func main() {
	fmt.Println("Maps")

	usuario := map[string]string{
		"nome":      "Pedro",
		"sobrenome": "Silva",
	}

	fmt.Println(usuario["nome"])

	// map dentro de map
	usuario2 := map[string]map[string]string{
		"nome": {
			"primeiro": "Rafael",
			"ultimo":   "jardim",
		},
	}
	fmt.Println(usuario2)

	delete(usuario2, "nome")
}
