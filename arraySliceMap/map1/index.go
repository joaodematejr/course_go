package main

import "fmt"

func main() {
	var aprovados map[int]string

	aprovados = make(map[int]string)

	aprovados[00000000000] = "Maria"
	aprovados[11111111111] = "Pedro"
	aprovados[22222222222] = "Ana"
	aprovados[33333333333] = "João"
	aprovados[44444444444] = "José"
	aprovados[55555555555] = "Paulo"
	aprovados[66666666666] = "Ricardo"
	aprovados[77777777777] = "Maria"
	aprovados[88888888888] = "Pedro"
	aprovados[99999999999] = "Ana"
	//fmt.Println(aprovados)

	for cpf, nome := range aprovados {
		fmt.Printf("%s (CPF: %d)\n", nome, cpf)
	}

	fmt.Println(aprovados[33333333333])

	delete(aprovados, 33333333333)

	fmt.Println(aprovados[33333333333])
}
