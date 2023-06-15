package main

import "fmt"

func main() {
	var num int

	fmt.Print("Digite um número: ")
	fmt.Scan(&num)

	fmt.Println("Sequência de Fibonacci:")
	for i := 0; i <= num; i++ {
		fmt.Println(fibonacci(i))
	}
}

func fibonacci(n int) int {
	if n <= 1 {
		return n
	}

	return fibonacci(n-1) + fibonacci(n-2)
}
