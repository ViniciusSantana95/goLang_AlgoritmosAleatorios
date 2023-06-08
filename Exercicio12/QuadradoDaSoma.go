package main

import (
	"fmt"
)

func main() {

	var num1 int
	var num2 int
	var num3 int
	var soma int
	var quadrado int

	println("Informe o primeiro valor inteiro")
	fmt.Scan(&num1)

	println("Informe o segundo valor inteiro")
	fmt.Scan(&num2)
	println("Informe o terceiro valor inteiro")
	fmt.Scan(&num3)

	soma = num1 + num2 + num3
	quadrado = soma * soma

	fmt.Println("O quadrado da soma dos 3 valores é: ", quadrado)

}
