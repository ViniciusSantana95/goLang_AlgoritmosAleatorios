package main

import "fmt"

func main() {

	var salario float32
	var reajuste float32
	var salarioReajuste float32

	fmt.Print("Informe o salario: ")
	fmt.Scan(&salario)
	fmt.Print("Informe o valor do reajuste: ")
	fmt.Scan(&reajuste)

	salarioReajuste = salario - reajuste

	fmt.Printf("O salario calculado é de : %.2f", salarioReajuste)

}
