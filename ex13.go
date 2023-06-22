package main

import "fmt"

//Crie um Array de inteiros com 7 elementos. Solicite ao usuário que informe um número que será adicionado ao primeiro
//e ao último elemento do Array. Imprima o Array resultante.

func main() {
	arr := [7]int{}
	fmt.Print("Digite o primeiro elemento do array: ")
	fmt.Scan(&arr[0])
	fmt.Print("Digite o último elemento do array: ")
	fmt.Scan(&arr[6])
	for i := 1; i < len(arr)-1; i++ {
		arr[i] = arr[0] + i
	}
	fmt.Print(arr)
}
