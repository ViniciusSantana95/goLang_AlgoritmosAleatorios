package main

import (
	"fmt"
	"strings"
)

func breakLines(frase string, colunas int) string {
	palavras := strings.Fields(frase)
	var linhas []string
	linhaAtual := ""

	for _, palavra := range palavras {
		if len(linhaAtual)+len(palavra) <= colunas {
			linhaAtual += palavra + " "
		} else {
			linhas = append(linhas, linhaAtual)
			linhaAtual = palavra + " "
		}
	}

	linhas = append(linhas, linhaAtual)
	return strings.Join(linhas, "\n")
}

func main() {
	var frase string
	var colunas int

	fmt.Print("Digite uma frase: ")
	fmt.Scan(&frase)

	fmt.Print("Digite a quantidade de colunas: ")
	fmt.Scan(&colunas)

	fraseQuebrada := breakLines(frase, colunas)
	fmt.Println(fraseQuebrada)
}
