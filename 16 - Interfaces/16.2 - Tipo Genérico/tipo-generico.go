package main

import "fmt"

func generica(interf interface{}) {
	fmt.Println(interf)
}

func main() {
	//Com a construção do da interface genérica, não temos restrição de tipo para receber
	generica("String")
	generica(1)
	generica(true)

	//Um exemplo disso é a função Print
}
