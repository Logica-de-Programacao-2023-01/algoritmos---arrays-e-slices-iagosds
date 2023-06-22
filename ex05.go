package main

import "fmt"

//Crie um Array bidimensional de inteiros com 3 linhas e 2 colunas. Solicite ao usuário que informe os valores de
//cada elemento da matriz. Em seguida, imprima a matriz resultante.

func main() {
	arr := [3][2]int{}
	for i := range arr {
		for c := range arr[i] {
			fmt.Printf("Digite o valor da %dª linha e %dª coluna: ", i+1, c+1)
			fmt.Scan(&arr[i][c])
		}
	}
	for i := range arr {
		fmt.Println(arr[i])
	}
}
