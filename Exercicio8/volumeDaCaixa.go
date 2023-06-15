package main

import "fmt"

func main() {
	var comprimento, largura, altura float64

	fmt.Print("Digite o comprimento da caixa: ")
	fmt.Scan(&comprimento)

	fmt.Print("Digite a largura da caixa: ")
	fmt.Scan(&largura)

	fmt.Print("Digite a altura da caixa: ")
	fmt.Scan(&altura)

	volume := comprimento * largura * altura

	fmt.Println("O volume da caixa é:", volume)
}
