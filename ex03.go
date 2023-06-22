package main

import "fmt"

//Crie um Array de floats com 4 elementos e calcule o produto dos valores armazenados no Array

func main() {
	num := []float64{2.1, 3.2, 4.1, 5.4}
	var prod float64 = 1
	for i := range num {
		prod *= num[i]
	}
	fmt.Printf("Produto: %.2f", prod)
}
