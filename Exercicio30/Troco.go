package main

import (
	"fmt"
)

func calcularTroco(valorTotal, valorPago float64) map[float64]int {
	valorTroco := valorPago - valorTotal

	cedulas := []float64{100, 50, 10, 5, 1}
	moedas := []float64{0.5, 0.1, 0.05, 0.01}

	troco := make(map[float64]int)

	for _, cedula := range cedulas {
		if valorTroco >= cedula {
			quantidadeCedulas := int(valorTroco / cedula)
			troco[cedula] = quantidadeCedulas
			valorTroco -= float64(quantidadeCedulas) * cedula
		}
	}

	for _, moeda := range moedas {
		if valorTroco >= moeda {
			quantidadeMoedas := int(valorTroco / moeda)
			troco[moeda] = quantidadeMoedas
			valorTroco -= float64(quantidadeMoedas) * moeda
		}
	}

	return troco
}

func exibirTroco(troco map[float64]int) {
	fmt.Println("Troco:")
	for valor, quantidade := range troco {
		if valor >= 1 {
			fmt.Printf("%d cédula(s) de R$%.2f\n", quantidade, valor)
		} else {
			valorCentavos := int(valor * 100)
			fmt.Printf("%d moeda(s) de R$%.2f\n", quantidade, float64(valorCentavos)/100)
		}
	}
}

func main() {
	var valorTotal, valorPago float64

	fmt.Print("Informe o valor total a ser pago: ")
	fmt.Scan(&valorTotal)

	fmt.Print("Informe o valor efetivamente pago: ")
	fmt.Scan(&valorPago)

	troco := calcularTroco(valorTotal, valorPago)

	exibirTroco(troco)
}
