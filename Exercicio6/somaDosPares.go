package main

import "fmt"

func main() {
	var start, end int

	fmt.Print("Digite o número inicial do intervalo: ")
	fmt.Scan(&start)

	fmt.Print("Digite o número final do intervalo: ")
	fmt.Scan(&end)

	sum := 0

	for i := start; i <= end; i++ {
		if i%2 == 0 {
			sum += i
		}
	}

	fmt.Println("A soma dos números pares no intervalo é:", sum)
}
