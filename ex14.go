package main

import "fmt"

//Crie um Slice de inteiros com tamanho 8 e solicite ao usuário que informe dois índices de elementos que deverão ser
//trocados de posição. Imprima o Slice resultante

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7}
	var ind, ind2, aux int
	fmt.Print("Indice 1: ")
	fmt.Scan(&ind)
	fmt.Print("Indice 2: ")
	fmt.Scan(&ind2)
	aux = arr[ind]
	arr[ind] = arr[ind2]
	arr[ind2] = aux
	fmt.Print(arr)

}
