package main

import "fmt"

func main() {

	retorno := func(texto string) string {
		return fmt.Sprintf("Olá %s", texto)
	}("universo")

	fmt.Println(retorno)
}
