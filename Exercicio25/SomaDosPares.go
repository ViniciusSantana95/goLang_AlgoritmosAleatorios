package main

import (
	"fmt"
	"strings"
)

func main() {

	continuar := "SIM"
	var soma float32
	var media float32
	var numeros []float32
	var valor float32
	for continuar == "SIM" {
		fmt.Println("Informe numeros, para saber a média deles")
		fmt.Scan(&valor)
		numeros = append(numeros, valor)

		fmt.Println("Informe 'sim' para continuar, ou informe outro valor para sair")
		fmt.Scan(&continuar)
		continuar = strings.ToUpper(continuar)
	}
	for i := 0; i < len(numeros); i++ {
		soma += numeros[i]
		if i+1 == len(numeros) {
			media = soma / (float32(i + 1))
			fmt.Printf("Foi informado %d numero(s), com a média de: %f \n", i+1, media)
		}
	}
}
