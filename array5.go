package main

import "fmt"

func main() {
	var mat [3][4]int
	for linha := range mat {
		for coluna := range mat[linha] {
			mat[linha][coluna] = linha + coluna
		}
	}
	for i := 0; i < len(mat); i++ {
		fmt.Println(mat[i])
	}
}
