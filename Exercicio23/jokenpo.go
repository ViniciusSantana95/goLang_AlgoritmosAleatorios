package main

import (
	"fmt"
)

func main() {
	var jogador1, jogador2 string

	// Leitura das jogadas dos jogadores
	fmt.Println("Jogador 1, escolha: Pedra, Papel ou Tesoura")
	fmt.Scan(&jogador1)

	fmt.Println("Jogador 2, escolha: Pedra, Papel ou Tesoura")
	fmt.Scan(&jogador2)

	// Verificação das jogadas e determinação do resultado
	switch jogador1 {
	case "Pedra":
		switch jogador2 {
		case "Pedra":
			fmt.Println("Empate!")
		case "Papel":
			fmt.Println("Jogador 2 ganhou!")
		case "Tesoura":
			fmt.Println("Jogador 1 ganhou!")
		default:
			fmt.Println("Jogada inválida do Jogador 2")
		}
	case "Papel":
		switch jogador2 {
		case "Pedra":
			fmt.Println("Jogador 1 ganhou!")
		case "Papel":
			fmt.Println("Empate!")
		case "Tesoura":
			fmt.Println("Jogador 2 ganhou!")
		default:
			fmt.Println("Jogada inválida do Jogador 2")
		}
	case "Tesoura":
		switch jogador2 {
		case "Pedra":
			fmt.Println("Jogador 2 ganhou!")
		case "Papel":
			fmt.Println("Jogador 1 ganhou!")
		case "Tesoura":
			fmt.Println("Empate!")
		default:
			fmt.Println("Jogada inválida do Jogador 2")
		}
	default:
		fmt.Println("Jogada inválida do Jogador 1")
	}
}
