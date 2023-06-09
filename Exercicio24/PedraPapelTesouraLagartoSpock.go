package main

import (
	"fmt"
	"strings"
)

func main() {
	jogarNovamente := "SIM"
	for jogarNovamente == "SIM" {
		jogador1 := 6
		jogador2 := 6

		fmt.Println("Escolha de 1 - 5 ")
		fmt.Println("1 - Tesoura")
		fmt.Println("2 - Papel")
		fmt.Println("3 - Pedra")
		fmt.Println("4 - Lagarto")
		fmt.Println("5 - Spock")
		for jogador1 <= 0 || jogador1 > 5 {
			fmt.Println("Jogador 1, faça sua jogada de 1 - 5")
			fmt.Scan(&jogador1)
		}

		obterOpcao(jogador1)
		for jogador2 <= 0 || jogador2 > 5 {
			fmt.Println("jogador 2, faça sua jogada de 1 - 5")
			fmt.Scan(&jogador2)
		}

		obterOpcao(jogador2)

		validarJogada(jogador1, jogador2)

		fmt.Println("Deseja jogar novamente ? Digite 'sim' para continuar, outra tecla para sair")
		fmt.Scan(&jogarNovamente)
		jogarNovamente = strings.ToUpper(jogarNovamente)
	}
}

func validarJogada(jogada1 int, jogada2 int) {

	if jogada1 == 1 && jogada2 == 2 || jogada2 == 1 && jogada1 == 2 {
		fmt.Println("Tesoura corta Papel!")
	} else if jogada1 == 2 && jogada2 == 3 || jogada2 == 2 && jogada1 == 3 {
		fmt.Println("Papel cobre Pedra!")
	} else if jogada1 == 3 && jogada2 == 4 || jogada2 == 3 && jogada1 == 4 {
		fmt.Println("Pedra esmaga Lagarto!")
	} else if jogada1 == 4 && jogada2 == 5 || jogada2 == 4 && jogada1 == 5 {
		fmt.Println("Lagarto envenena Spock!")
	} else if jogada1 == 5 && jogada2 == 1 || jogada2 == 5 && jogada1 == 1 {
		fmt.Println("Spock esmaga (ou derrete) tesoura!")
	} else if jogada1 == 1 && jogada2 == 4 || jogada2 == 1 && jogada1 == 4 {
		fmt.Println("Tesoura decapita lagarto!")
	} else if jogada1 == 4 && jogada2 == 2 || jogada2 == 4 && jogada1 == 2 {
		fmt.Println("Lagarto come papel!")
	} else if jogada1 == 2 && jogada2 == 5 || jogada2 == 2 && jogada1 == 5 {
		fmt.Println("Papel refuta Spock!")
	} else if jogada1 == 5 && jogada2 == 3 || jogada2 == 5 && jogada1 == 3 {
		fmt.Println("Spock vaporiza pedra!")
	} else if jogada1 == 3 && jogada2 == 1 || jogada2 == 3 && jogada1 == 1 {
		fmt.Println("Pedra amassa tesoura!")
	} else {
		fmt.Println("Empate!")
	}
}

func obterOpcao(valor int) {
	switch valor {
	case 1:
		fmt.Println("A jogada escolhida foi Tesoura!")
	case 2:
		fmt.Println("A jogada escolhida foi Papel!")
	case 3:
		fmt.Println("A jogada escolhida foi Pedra!")
	case 4:
		fmt.Println("A jogada escolhida foi Lagarto!")
	case 5:
		fmt.Println("A jogada escolhida foi Spock!")
	default:
		break
	}
}
