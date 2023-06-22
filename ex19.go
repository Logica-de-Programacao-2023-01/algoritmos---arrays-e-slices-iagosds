package main

import "fmt"

//Faça um programa que leia dois arrays de inteiros de tamanho n e gere um terceiro array que seja a soma dos dois primeiros

func main() {
	arr := []int{1, 2, 3, 4, 5}
	arr2 := []int{6, 7, 8, 9, 10}
	res := []int{}
	for i := range arr {
		res = append(res, arr[i])
	}
	for i := range arr2 {
		res = append(res, arr2[i])
	}
	fmt.Println(arr)
	fmt.Println(arr2)
	fmt.Println(res)
}
