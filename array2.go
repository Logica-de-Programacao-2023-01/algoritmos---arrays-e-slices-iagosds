package main

import "fmt"

func main() {
	var nome []string
	nome = append(nome, "joao", "maria", "josé")
	fmt.Println("O array possui ", len(nome), " elementos")
	fmt.Println(nome)
}
