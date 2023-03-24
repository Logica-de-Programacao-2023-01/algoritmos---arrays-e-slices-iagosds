package main

import "fmt"

func main() {
	var n int
	fmt.Println("Quantos elementos terão no string? ")
	fmt.Scan(&n)
	var sl [n]int
	for i := 0; i < n; i++ {
		fmt.Print("Escolha o ", i+1, "º valor do slice: ")
		fmt.Scanln(&sl[i])
	}
	fmt.Print(sl)
}
