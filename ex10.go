package main

import "fmt"

//Crie um Slice de inteiros com tamanho 10 e imprima o valor mínimo e o valor máximo armazenados no Slice

func main() {
	num := []int{2, 4, 1, 15, 3, 5}
	var menor, maior int = num[0], num[0]
	for i := range num {
		if num[i] < menor {
			menor = num[i]
		} else if num[i] > maior {
			maior = num[i]
		}
	}
	fmt.Println("Max: ", maior)
	fmt.Print("Min: ", menor)
}
