package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	var nome string

	println("Informe um nome !")
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		nome = scanner.Text()
	}

	if nome == "" {
		fmt.Println("Um para você, um para mim!")
	} else {
		fmt.Printf("Um para %s, um para mim!\n", nome)
	}

}
