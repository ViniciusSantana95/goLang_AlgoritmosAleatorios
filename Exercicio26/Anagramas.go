package main

import (
	"fmt"
	"strings"
)

func main() {

	var palavra1 string
	var palavra2 string

	fmt.Println("Informe a primeira palavra")
	fmt.Scan(&palavra1)

	palavra1 = strings.ToUpper(palavra1)

	fmt.Println("Informe a segunda palavra")
	fmt.Scan(&palavra2)
	palavra2 = strings.ToUpper(palavra2)

	fmt.Println(anagrama(palavra1, palavra2))

}

func anagrama(palavra string, palavra2 string) string {
	var validacao []bool
	for i := 0; i < len(palavra); i++ {
		validacao = append(validacao, false)
	}
	if len(palavra) != len(palavra2) {
		return "Não é anagrama, as palavras possuem quantidade de letras diferentes!"
	}

	for i := 0; i < len(palavra); i++ {

		for j := 0; j < len(palavra2); j++ {
			if validacao[j] == true {
				//fmt.Printf("Posição %d já foi verificada \n", j)
				//a linha acima foi feita durante a fase de testes, para verificação se a letra ja havia sido verificada, para n ser usada novamente
				continue
			} else if palavra[i] == palavra2[j] {
				validacao[j] = true
				fmt.Printf("letra %c encontrada na posição %d %d \n", palavra[i], i+1, j+1)
				break
			} else if j == len(palavra)-1 && validacao[j] == false {
				return "Não é anagrama"
			} //else {
			//fmt.Printf("letra %c não encontrada na posição %d %d \n", palavra[i], i+1, j+1)
			//linhas usadas para informar que a letra não foi encontrada, porém para melhor visibilidade para o usuario, foi removida
			//}
		}
	}
	return "É anagrama"
}
