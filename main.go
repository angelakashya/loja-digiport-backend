package main

import (
	"fmt"
)

func main() {
	nome := ""
	fmt.Scan("Digite o nome do produto desejado", &nome)
	produtosFiltrados := buscaPorNome(nome)

	fmt.Println(produtosFiltrados)

}
