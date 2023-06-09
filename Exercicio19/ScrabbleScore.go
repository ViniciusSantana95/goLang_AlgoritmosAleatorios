package main

import (
	"fmt"
	"strings"
)

func main() {

	var palavra string
	var pontuacao int
	var letra rune
	fmt.Println("Informe uma palavra para saber seu score!")
	fmt.Println("Caso informe mais que uma palavra, apenas a primeira será considerada")
	fmt.Scan(&palavra)

	palavra = strings.ToLower(palavra)

	for i := 0; i < len(palavra); i++ {

		letra = rune(palavra[i])

		switch letra {
		case 'a', 'e', 'i', 'o', 'u', 'l', 'n', 'r', 's', 't':
			pontuacao += 1
		case 'd', 'g':
			pontuacao += 2
		case 'b', 'c', 'm', 'p':
			pontuacao += 3
		case 'f', 'h', 'v', 'w', 'y':
			pontuacao += 4
		case 'k':
			pontuacao += 5
		case 'j', 'x':
			pontuacao += 8
		case 'q', 'z':
			pontuacao += 10
		default:
			pontuacao += 0
		}

	}

	fmt.Printf("A pontuação da palavra %s foi de %d \n", palavra, pontuacao)

}
