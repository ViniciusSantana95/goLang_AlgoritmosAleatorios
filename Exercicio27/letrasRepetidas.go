package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Digite um texto: ")
	texto, _ := reader.ReadString('\n')

	texto = strings.TrimSpace(texto)

	letra := encontrarPrimeiraLetraNaoRepetida(texto)

	if letra != "" {
		fmt.Println("A primeira letra não repetida é:", letra)
	} else {
		fmt.Println("Não existem letras que não se repetem no texto informado.")
	}
}

func encontrarPrimeiraLetraNaoRepetida(texto string) string {
	frequencia := make(map[rune]int)

	for _, char := range texto {
		frequencia[char]++
	}

	for _, char := range texto {
		if frequencia[char] == 1 {
			return string(char)
		}
	}

	return ""
}
