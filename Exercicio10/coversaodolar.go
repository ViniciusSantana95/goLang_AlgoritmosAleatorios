package main

import "fmt"

func main() {
	var dolar float32
	var real float32
	var coversao float32

	fmt.Print("Informe o valor atual do dolar: ")
	fmt.Scan(&dolar)
	fmt.Print("Informe o valor atual do real: ")
	fmt.Scan(&real)

	coversao = real / dolar

	fmt.Printf("O resultado é: %.2f", coversao)
}
