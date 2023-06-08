package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Informe uma palavra para verificar se é palíndromo:")
	palindromo, _ := reader.ReadString('\n')
	palindromo = strings.TrimSpace(palindromo)

	if ePalindromo(palindromo) {
		fmt.Println("É palíndromo")
	} else {
		fmt.Println("Não é palíndromo")
	}
}

func ePalindromo(palavra string) bool {
	palavra = strings.ToLower(strings.ReplaceAll(palavra, " ", ""))
	var novaPalavra strings.Builder
	for i := len(palavra) - 1; i >= 0; i-- {
		novaPalavra.WriteByte(palavra[i])
	}
	return palavra == novaPalavra.String()
}
