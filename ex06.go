package main

import "fmt"

//Crie um Array de inteiros com 3 elementos. Em seguida, solicite ao usuário que informe um valor e verifique se esse
//valor está presente no Array. Imprima o resultado da busca

func main() {
	num := [3]int{10, 20, 30}
	valor := 0
	pres := false
	fmt.Print("Digite um valor: ")
	fmt.Scan(&valor)
	for i := range num {
		if valor == num[i] {
			fmt.Printf("Valor presente na %dª posição.", i+1)
			pres = true
		}
	}
	if pres == false {
		fmt.Println("Número não encontrado")
	}
}
