package main

import "fmt"

//Crie um Array de inteiros com 5 elementos. Em seguida, crie um novo Slice que contenha apenas os elementos do Array
//que são múltiplos de 3. Imprima o novo Slice.

func main() {
	arr := [5]int{2, 3, 5, 6, 9}
	mult := []int{}
	for i := range arr {
		if arr[i]%3 == 0 {
			mult = append(mult, arr[i])
		}
	}
	fmt.Print(mult)
}
