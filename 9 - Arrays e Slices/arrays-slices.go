package main

import "fmt"

func main() {

	fmt.Println("Ponteiros")

	var array1 [5]int // [0 0 0 0]
	fmt.Println(array1)

	array2 := [5]string{"0", "1", "2", "3", "4"}
	fmt.Println(array2)

	array3 := [...]int{1, 2, 3, 4, 5}
	fmt.Println(array3)

	// Slice
	slice := []int{}
	fmt.Println(slice)

	slice = append(slice, 18)
	slice2 := array2[1:3]
	fmt.Println(slice2)

	// Arrays Internos
	fmt.Println("--------")
	slice3 := make([]float32, 10, 11)

	slice3 = append(slice3, 5)
	slice3 = append(slice3, 6)

	fmt.Println(slice3)
	fmt.Println(len(slice3))
	fmt.Println(cap(slice3))
	// Ele dobra a capacidade quando tentamos estourar
}
