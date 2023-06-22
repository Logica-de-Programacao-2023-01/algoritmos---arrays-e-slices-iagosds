package main

import "fmt"

//Crie um Array de inteiros com 10 elementos. Calcule e imprima a soma dos elementos nas posições pares do Array

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	soma := 0
	for i := range arr {
		if i%2 == 0 {
			soma += arr[i]
		}
	}
	fmt.Print(soma)
}
