package main

import "fmt"

//Crie um Array bidimensional de inteiros com 2 linhas e 3 colunas. Em seguida, solicite ao usuário que informe um
//índice de linha e outro de coluna e imprima o valor armazenado nessa posição da matriz

func main() {
	mat := [2][3]int{}
	var linha, coluna int
	fmt.Print("Linha: ")
	fmt.Scan(&linha)
	fmt.Print("Coluna: ")
	fmt.Scan(&coluna)
	for i := range mat {
		for c := range mat[i] {
			mat[i][c] = i + c
		}
	}
	fmt.Print("Elemento: ", mat[linha][coluna])
}
