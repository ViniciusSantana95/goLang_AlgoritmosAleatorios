package main

import (
	"fmt"
	"strings"
)

func main() {
	var romano string
	fmt.Println("Informe um numeral romano para conversão")
	fmt.Scan(&romano)

	romano = strings.ToUpper(romano)
	numero := ConverterRomanoParaDecimal(romano)

	if numero > 0 {
		fmt.Printf("O número romano informado é: %d\n", numero)
	} else {
		fmt.Println("Número Romano inválido!")
	}
}

func ConverterRomanoParaDecimal(romano string) int {
	numero := 0
	var letra rune

	for i := 0; i < len(romano); i++ {
		letra = rune(romano[i])

		switch letra {
		case 'I':
			if i > 2 && romano[i-1] == 'I' && romano[i-2] == 'I' && romano[i-3] == 'I' {
				return -1
			} else {
				numero += 1
			}
		case 'V':
			if i > 0 && romano[i-1] == 'I' {
				numero += 3
			} else if romano[i-1] == 'V' {
				return -1
			} else {
				numero += 5
			}
		case 'X':
			if i > 0 && romano[i-1] == 'I' {
				numero += 8
			} else if i > 2 && romano[i-1] == 'X' && romano[i-2] == 'X' && romano[i-3] == 'X' {
				return -1
			} else {
				numero += 10
			}
		case 'L':
			if i > 0 && romano[i-1] == 'X' {
				numero += 30
			} else if i > 0 && romano[i-1] == 'L' {
				return -1
			} else {
				numero += 50
			}
		case 'C':
			if i > 0 && romano[i-1] == 'X' {
				numero += 80
			} else if i > 2 && romano[i-1] == 'C' && romano[i-2] == 'C' && romano[i-3] == 'C' {
				return -1
			} else {
				numero += 100
			}
		case 'D':
			if i > 0 && romano[i-1] == 'C' {
				numero += 300
			} else if i > 0 && romano[i-1] == 'D' {
				return -1
			} else {
				numero += 500
			}
		case 'M':
			if i > 0 && romano[i-1] == 'C' {
				numero += 800
			} else if i > 2 && romano[i-1] == 'C' {
				numero += 900
			} else {
				numero += 1000
			}
		default:
			return -1
		}
	}

	return numero
}
