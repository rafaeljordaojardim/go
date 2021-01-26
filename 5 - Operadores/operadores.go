package main

import "fmt"

func main() {
	// Aritimeticos
	soma := 1 + 1
	sub := 1 - 1
	divisao := 10 / 4
	multiplicacao := 10 * 5
	restoDaDivisao := 10 % 2

	fmt.Println(soma, sub, divisao, multiplicacao, restoDaDivisao)

	var numero1 int16 = 10
	var numero2 int16 = 25

	soma2 := numero1 + numero2
	fmt.Println(soma2)

	// Operadores de atribuição
	var variavel1 string = "string"
	variavel2 := "string2"
	fmt.Println(variavel1, variavel2)

	// Operadores Relacionais
	fmt.Println(1 > 2)
	fmt.Println(1 >= 2)
	fmt.Println(1 == 2)
	fmt.Println(1 <= 2)
	fmt.Println(1 < 2)
	fmt.Println(1 != 2)

	// Operadores Logicos
	verdadeiro, falso := true, false
	fmt.Println(verdadeiro && falso)
	fmt.Println(verdadeiro || falso)
	fmt.Println(!falso)
	fmt.Println(!verdadeiro)

	// Operadores unarios
	numero := 10
	numero++
	numero += 15
	fmt.Println(numero)

	// Operador ternario
	// não tem
}
