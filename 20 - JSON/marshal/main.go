package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
)

type cachorro struct {
	Nome  string `json:"nome"`
	Raca  string `json:"raca"`
	Idade uint   `json:"idade"`
}

func main() {
	c := cachorro{"Rex", "Dálmata", 3}
	fmt.Println(c)

	cachorroEmJSON, erro := json.Marshal(c)
	if erro != nil {
		log.Fatal()
	}

	fmt.Println(cachorroEmJSON)

	fmt.Println(bytes.NewBuffer(cachorroEmJSON))
}
