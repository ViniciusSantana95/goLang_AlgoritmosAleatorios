package main

import "fmt"

func main() {
	var raio float32
	const pi = 3.14159265
	var area float32

	fmt.Print("Informe o valor do raio: ")
	fmt.Scan(&raio)

	area = pi * raio * raio

	fmt.Printf("A área do círculo é: %.2f", area)

}
