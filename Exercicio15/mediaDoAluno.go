package main

import "fmt"

func main() {
	var nota1, nota2, nota3, nota4 float64

	fmt.Print("Digite a primeira nota: ")
	fmt.Scan(&nota1)

	fmt.Print("Digite a segunda nota: ")
	fmt.Scan(&nota2)

	fmt.Print("Digite a terceira nota: ")
	fmt.Scan(&nota3)

	fmt.Print("Digite a quarta nota: ")
	fmt.Scan(&nota4)

	media := (nota1 + nota2 + nota3 + nota4) / 4

	fmt.Println("Média obtida:", media)

	if media >= 50 {
		fmt.Println("Aprovado")
	} else {
		fmt.Println("Reprovado")
	}
}
