package main

import "fmt"

func main() {
	canal := make(chan string, 2)

	canal <- "ola mundo"

	mensagem := <-canal
	fmt.Println(mensagem)
}
