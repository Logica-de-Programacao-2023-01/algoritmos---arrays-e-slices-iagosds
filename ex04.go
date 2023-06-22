package main

import "fmt"

//Crie um Slice de inteiros e solicite ao usuário que informe o tamanho do Slice e os valores dos elementos.
//Em seguida, imprima o Slice e a soma dos valores armazenados.

func main() {
	tamanho, soma := 0, 0
	fmt.Print("Qual o tamanho do slice? ")
	fmt.Scan(&tamanho)
	num := make([]int, tamanho)
	for i := range num {
		fmt.Printf("%dº número: \t", i+1)
		fmt.Scan(&num[i])
		soma += num[i]
	}
	fmt.Println()
	for i := range num {
		fmt.Printf("%dº valor: %d\t", i+1, num[i])
	}
	fmt.Printf("\nSoma: %d", soma)

}
