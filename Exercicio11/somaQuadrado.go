package exercicio11

import (
	"fmt"
	"math"
)

func SomaQuadrados() {
	var num1, num2, num3 int

	fmt.Println("Digite o primeiro valor:")
	fmt.Scanln(&num1)

	fmt.Println("Digite o segundo valor:")
	fmt.Scanln(&num2)

	fmt.Println("Digite o terceiro valor:")
	fmt.Scanln(&num3)

	resultado := math.Pow(float64(num1), 2) + math.Pow(float64(num2), 2) + math.Pow(float64(num3), 2)

	fmt.Printf("A soma dos quadrados dos valores é: %.2f\n", resultado)
}
