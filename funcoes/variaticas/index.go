package main

import "fmt"

func media(numeros ...float64) float64 {
	total := 0.0
	for _, numero := range numeros {
		total += numero
	}
	return total / float64(len(numeros))
}

func main() {
	fmt.Println("Média: %.5f", media(7.7, 8.1, 5.9, 9.9))
}
