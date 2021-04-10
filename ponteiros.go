package main

import "fmt"

func soma(numeros []int) int {
	var soma int

	fmt.Println(&numeros[1])
	for i:= range numeros {
		soma = soma + numeros[i]
	}

	return soma
}

func changePointer(num *int) {
	fmt.Println(num)
	fmt.Println(*num)

	*num = 101
}

func main() {
	num := 89

	changePointer(&num)

	fmt.Println(num)
}