package main

import "fmt"

//Crie um Array de inteiros com 3 elementos e imprima a soma dos valores armazenados no Array.

func main() {
	num := []int{1, 2, 3}
	soma := 0
	for i := range num {
		fmt.Printf("%dº Valor: %d\t", i+1, num[i])
		soma += num[i]
	}
	fmt.Printf("\nSoma: %d", soma)
}
