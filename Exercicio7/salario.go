package main

import (
	"fmt"
)

func main() {

	var salarioHora float64
	var horaTrabMes int
	var perDesInss float64
	var salarioBruto float64
	var salarioLiquido float64

	fmt.Print("Informe o valor da Hora: ")
	fmt.Scan(&salarioHora)
	fmt.Print("Informe o numero de horas trabalhadas: ")
	fmt.Scan(&horaTrabMes)
	fmt.Print("Informe o desconto de INSS: ")
	fmt.Scan(&perDesInss)

	salarioBruto = salarioHora * float64(horaTrabMes)
	salarioLiquido = salarioBruto - (salarioBruto * (perDesInss / 100))

	fmt.Printf("O salário líquido é de: %.2f", salarioLiquido)

}
