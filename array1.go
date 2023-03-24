package main

import "fmt"

func main() {
	var num = [5]int{1, 2, 3, 4, 5}
	for i := 0; i < len(num); i++ {
		fmt.Println(num[i])
	}
}
