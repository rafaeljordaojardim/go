package main

import "fmt"

func main() {
	var variavel1 string = "Variável 1"
	variavel2 := "Variavel 2"
	fmt.Println(variavel1, variavel2)

	var (
		var3 string = "var3"
		var4 string = "var4"
	)

	fmt.Println(var3, var4)

	var5, var6 := "var5", "var6"
	fmt.Println(var5, var6)

	const constant string = "Constante"

	fmt.Println(constant)

	var5, var6 = var6, var5
}
