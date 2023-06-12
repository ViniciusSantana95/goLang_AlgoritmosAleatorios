package exercicio16

import "fmt"

func SomandoNaturais() {
	var inicio, fim, minimo, soma int

	fmt.Print("Digite o valor inicial do intervalo: ")
	fmt.Scanln(&inicio)
	fmt.Print("Digite o valor final do intervalo: ")
	fmt.Scanln(&fim)

	minimo = fim - inicio
	if inicio >= 1 && fim >= inicio && minimo >= 100 {
		for i := inicio; i <= fim; i++ {
			soma += i
		}

		fmt.Printf("Intervalo informado: %d a %d\n", inicio, fim)
		fmt.Printf("Soma dos números no intervalo: %d\n", soma)
	} else {
		fmt.Println("Intervalo inválido! O intervalo deve conter 100 números naturais.")
	}
}
