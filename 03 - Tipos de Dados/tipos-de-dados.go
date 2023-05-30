package main

import (
	"errors"
	"fmt"
)

func main() {
	var numero int64 = 100000000
	fmt.Println(numero)

	var numero2 uint32 = 10000
	fmt.Println(numero2)

	// alias
	//INT32 = RUNE
	var numero3 rune = 123456
	fmt.Println(numero3)

	//BYTE = UINT8
	var numero4 byte = 123
	fmt.Println(numero4)

	var numeroReal1 float32 = 123.45
	fmt.Println(numeroReal1)

	var numeroReal2 float64 = 1230000000.45
	fmt.Println(numeroReal2)

	//  FIM NÚMEROS

	var str string = "Texto"
	fmt.Println(str)
	
	str2 := "Texto2"
	fmt.Println(str2)

	char := 'A'
	fmt.Println(char)

	// FIM STRING

	var texto int16
	fmt.Println(texto)

	var booleano1 bool
	fmt.Println(booleano1)

	//Usando o pacote errors para atribuir uma escrita no retorno do error
	var erro error = errors.New("erro interno")
	println(erro)
}