package main

import "fmt"

func main() {
	fmt.Println("Estrutura de controle")

	numero := 10

	if outroNumero := numero; outroNumero > 0 {
		fmt.Println("Numero maior que 0")
	} else if numero < -10 {
		fmt.Println("Número menor que -10")
	} else {
		fmt.Println("Entre 0 e -10")
	}
	
}