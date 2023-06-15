package main

import (
	"fmt"
	"sort"
)

func main() {
	numeros := make([]int, 12)

	for i := 0; i < len(numeros); i++ {
		fmt.Printf("Digite o elemento %d: ", i+1)
		fmt.Scan(&numeros[i])
	}

	sort.Sort(sort.Reverse(sort.IntSlice(numeros)))

	fmt.Println("Elementos em ordem decrescente:")
	for _, numero := range numeros {
		fmt.Println(numero)
	}
}
