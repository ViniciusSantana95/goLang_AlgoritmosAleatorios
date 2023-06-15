package main

import (
	"fmt"
)

func main() {
	var start, end int

	fmt.Print("Digite o número inicial do intervalo: ")
	fmt.Scan(&start)

	fmt.Print("Digite o número final do intervalo: ")
	fmt.Scan(&end)

	for i := start; i <= end; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println("FizzBuzz")
		} else if i%3 == 0 {
			fmt.Println("Fizz")
		} else if i%5 == 0 {
			fmt.Println("Buzz")
		} else {
			fmt.Println(i)
		}
	}
}
