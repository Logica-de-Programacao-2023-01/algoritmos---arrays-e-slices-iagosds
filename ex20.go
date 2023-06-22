package main

import "fmt"

//Crie um programa que leia um array de inteiros e verifique se ele está ordenado em ordem crescente.

func main() {
	arr := []int{1, 0, 3, 4, 5}
	cres := true
	for i := 0; i < len(arr)-1; i++ {
		if arr[i] > arr[i+1] {
			cres = false
			break
		}
	}
	if len(arr) == 1 {
		cres = true
	}
	if cres {
		fmt.Println("Array em ordem crescente")
	} else {
		fmt.Println("Array desordenado")
	}
}
