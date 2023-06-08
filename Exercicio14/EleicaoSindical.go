package main

import "fmt"

func main() {

	var votosCandidato1 int
	var votosCandidato2 int
	var votosCandidato3 int
	var votosNulos int
	var votosBrancos int
	var totalEleitores int
	var votosValidos int
	var porcentagemValidos float32
	var percentualCandidato1 float32
	var percentualCandidato2 float32
	var percentualCandidato3 float32
	var percentualNulos float32
	var percentualBrancos float32

	println("Informe a quantidade de votos para o candidato 1")
	fmt.Scan(&votosCandidato1)
	println("Informe a quantidade de votos para o candidato 2")
	fmt.Scan(&votosCandidato2)
	println("Informe a quantidade de votos para o candidato 3")
	fmt.Scan(&votosCandidato3)
	println("Informe a quantidade de votos nulos")
	fmt.Scan(&votosNulos)
	println("Informa a quantidade de votos brancos")
	fmt.Scan(&votosBrancos)

	totalEleitores = votosCandidato1 + votosCandidato2 + votosCandidato3 + votosNulos + votosBrancos
	fmt.Println("A quantidade total de votos foi de: ", totalEleitores)

	votosValidos = votosCandidato1 + votosCandidato2 + votosCandidato3

	porcentagemValidos = (float32(totalEleitores) / 100) * float32(votosValidos)

	percentualCandidato1 = (float32(totalEleitores) / 100) * float32(votosCandidato1)

	percentualCandidato2 = (float32(totalEleitores) / 100) * float32(votosCandidato2)

	percentualCandidato3 = (float32(totalEleitores) / 100) * float32(votosCandidato3)

	percentualNulos = (float32(totalEleitores) / 100) * float32(votosNulos)

	percentualBrancos = (float32(totalEleitores) / 100) * float32(votosBrancos)

	fmt.Printf("A porcentagem de votos validos, em relação a quantidade de eleitores foi de: %.2f \n", porcentagemValidos)
	fmt.Printf("A porcentagem de votos do candidato 1, rem relação a quantidade de eleitores foi de: %.2f \n", percentualCandidato1)
	fmt.Printf("A porcentagem de votos do candidato 2, rem relação a quantidade de eleitores foi de: %.2f \n", percentualCandidato2)
	fmt.Printf("A porcentagem de votos do candidato 3, rem relação a quantidade de eleitores foi de: %.2f \n", percentualCandidato3)
	fmt.Printf("A porcentagem de votos nulos, em relação a quantidade de eleitores foi de: %.2f \n", percentualNulos)
	fmt.Printf("A porcentagem de votos brancos, em relação a quantidade de eleitores foi de: %.2f \n", percentualBrancos)

}
