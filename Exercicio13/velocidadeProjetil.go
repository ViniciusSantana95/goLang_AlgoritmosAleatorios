package exercicio13

import "fmt"

func VelocidadeProjetil() {
	var distancia float64
	var tempo float64

	fmt.Print("Digite a distância percorrida : ")
	fmt.Scanln(&distancia)

	fmt.Print("Digite o tempo decorrido : ")
	fmt.Scanln(&tempo)

	velocidade := (distancia * 1000) / (tempo * 60)

	fmt.Printf("A velocidade do projétil é de %.2f metros por segundo.\n", velocidade)
}
