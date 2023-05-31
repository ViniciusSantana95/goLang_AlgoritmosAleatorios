package main

import "fmt"

func main() {

	var nome string

	println("Informe um nome !")
	fmt.Scan(&nome)

	if nome == "" {
		fmt.Println("Um para você, um para mim!")
	} else {
		fmt.Printf("Um para %s, um para mim!\n", nome)
	}

}
