package main

import "fmt"

//Crie um Array de inteiros com 10 elementos. Crie um novo Slice que contenha apenas os elementos pares do Array
//original. Imprima o novo Slice

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	novo := []int{}
	for i := range arr {
		if arr[i]%2 == 0 {
			novo = append(novo, arr[i])
		}
	}
	fmt.Print(novo)
}
