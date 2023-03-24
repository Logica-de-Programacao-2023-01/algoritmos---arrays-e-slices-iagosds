package main

import (
	"fmt"
)

func main() {
	var num [6]float64
	var s float64
	for i := 0; i < len(num); i++ {
		fmt.Print("Escreva um número real: ")
		fmt.Scanln(&num[i])
		s = (s + num[i])
	}
	m := s / float64(len(num))
	fmt.Print("A média dos números é : ")
	fmt.Printf("%.2f", m)
}
