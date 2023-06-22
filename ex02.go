package main

import "fmt"

//Crie um Slice de inteiros com os valores 1, 2, 3, 4 e 5. Remova o terceiro elemento e imprima o Slice resultante.

func main() {
	num := []int{1, 2, 3, 4, 5}
	for i := range num {
		fmt.Printf("%dº valor: %d\t", i+1, num[i])
	}
	num = append(num[:3], num[4:]...)
	for i := range num {
		fmt.Printf("%dº valor: %d\t", i+1, num[i])
	}
}
