package main

import "fmt"

//Crie um Array de floats com 6 elementos. Solicite ao usuário que informe um número que será adicionado em todas as
//posições do Array. Imprima o Array resultante

func main() {
	arr := [4]float64{}
	fmt.Print("Qual valor das posições do array? ")
	fmt.Scan(&arr[0])
	for i := 1; i < len(arr); i++ {
		arr[i] = arr[0]
	}
	fmt.Print("Array: ", arr)
}
