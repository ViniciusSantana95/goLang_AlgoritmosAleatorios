package main

import "fmt"

func main() {
	var cotacao, quantidadeDolar float64

	fmt.Print("Digite a cotação do dólar: ")
	fmt.Scan(&cotacao)

	fmt.Print("Digite a quantidade de dólares disponível: ")
	fmt.Scan(&quantidadeDolar)

	valorEmReal := cotacao * quantidadeDolar

	fmt.Printf("O valor em reais é R$ %.2f\n", valorEmReal)
}
