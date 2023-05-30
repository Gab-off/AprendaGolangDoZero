package main

import "fmt"

func main() {
	fmt.Println("Maps")

	usuario := map[string]string {
		"nome" : "Pedro",
		"sobrenome" : "Silva",
	}

	fmt.Println(usuario)

	usuario2 := map[string]map[string]string {
		"nome" : {
			"primeiro" : "Tony",
			"ultimo" : "Montana",
		},
		"curso": {
			"nome" : "Engenharia",
			"periodo" : "primeiro",
		},
	}

	delete(usuario2, "nome")
	fmt.Println(usuario2)

}