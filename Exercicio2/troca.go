package main

import "fmt"

func main() {
	var A, B int

	fmt.Print("Digite o valor de A: ")
	fmt.Scan(&A)

	fmt.Print("Digite o valor de B: ")
	fmt.Scan(&B)

	A, B = B, A

	fmt.Println("Após a troca:")
	fmt.Println("Valor de A:", A)
	fmt.Println("Valor de B:", B)
}
