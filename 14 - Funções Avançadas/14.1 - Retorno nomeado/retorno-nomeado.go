package main

import "fmt"

func calculosMatematicos(n1, n2 int) (soma int, sub int) {
	soma = n1 + n2
	sub = n1 - n2
	return
}

func main() {
	soma, subt := calculosMatematicos(10, 20)
	fmt.Println(soma, subt)
}
