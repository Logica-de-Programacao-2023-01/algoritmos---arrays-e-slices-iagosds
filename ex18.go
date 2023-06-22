package main

import "fmt"

//Escreva um programa que leia um número inteiro positivo n e gere um array com os n primeiros números primos

func primo(x int) bool {
	if x <= 1 {
		return false
	}
	for i := 2; i < x/2; i++ {
		if x%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	var n, i, c int = 0, 1, 0
	fmt.Print("Digite um número: ")
	fmt.Scan(&n)
	n--
	arr := []int{}
	for c <= n {
		if primo(i) {
			arr = append(arr, i)
			c++
		}
		i++
	}
	fmt.Print(arr)
}
