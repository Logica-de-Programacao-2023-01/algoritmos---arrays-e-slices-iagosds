package main

import "fmt"

//Crie um Slice de inteiros com o tamanho 5. Em seguida, solicite ao usuário que informe um número inteiro.
//Adicione esse número ao Slice apenas se ele não estiver presente. Imprima o Slice resultante

func main() {
	num := []int{10, 20, 30, 40, 50}
	valor := 0
	pres := false
	fmt.Print("Digite um valor: ")
	fmt.Scan(&valor)
	for i := range num {
		if valor == num[i] {
			pres = true
		}
	}
	if !pres {
		num = append(num, valor)
	}
	fmt.Print("Slice: ", num)
}
