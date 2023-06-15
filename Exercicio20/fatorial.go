package main

import "fmt"

func main() {
	var num int

	fmt.Print("Digite um número: ")
	fmt.Scan(&num)

	fatorial := 1

	for i := 1; i <= num; i++ {
		fatorial *= i
	}

	fmt.Printf("O fatorial de %d é %d\n", num, fatorial)
}
