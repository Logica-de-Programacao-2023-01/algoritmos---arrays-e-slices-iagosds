package main

import "fmt"

//Crie um Array de floats com 10 elementos. Crie um novo Slice que contenha apenas os elementos do Array que são
//maiores que 5. Imprima o novo Slice

func main() {
	arr := []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	novo := []float64{}
	for i := range arr {
		if arr[i] > 5 {
			novo = append(novo, arr[i])
		}
	}
	fmt.Print(novo)
}
