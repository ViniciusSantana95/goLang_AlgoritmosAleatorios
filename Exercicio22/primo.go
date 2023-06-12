package main

import (
	"fmt"
	"math"
)

func isPrime(n int) bool {
	if n <= 1 {
		return false
	}

	// Verificar divisibilidade por números até a raiz quadrada de n
	for i := 2; i <= int(math.Sqrt(float64(n))); i++ {
		if n%i == 0 {
			return false
		}
	}

	return true
}

func main() {
	var num int
	fmt.Print("Digite um número: ")
	fmt.Scan(&num)

	if isPrime(num) {
		fmt.Println(num, "é um número primo.")
	} else {
		fmt.Println(num, "não é um número primo.")
	}
}
