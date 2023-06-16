package main

import (
	"fmt"
	"strings"
)

func isPangram(frase string) bool {

	frase = strings.ToLower(frase)

	letterCount := make(map[rune]int)

	for _, char := range frase {
		if char >= 'a' && char <= 'z' {
			letterCount[char]++
		}
	}

	return len(letterCount) == 26
}

func main() {
	var frase string
	fmt.Print("Digite uma frase: ")
	fmt.Scanln(&frase)

	if isPangram(frase) {
		fmt.Println("A frase é um pangrama.")
	} else {
		fmt.Println("A frase não é um pangrama.")
	}
}
