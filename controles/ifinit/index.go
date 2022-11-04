package main

import "fmt"

func tipo(i interface{}) string {
	switch i.(type) {
	case int:
		return "Inteiro"
	case float64, float32:
		return "Real"
	case string:
		return "String"
	case func():
		return "Função"
	default:
		return "Tipo não encontrado"
	}
}

func main() {
	fmt.Println(tipo(2.3))
	fmt.Println(tipo(1))
	fmt.Println(tipo("Opa"))
	fmt.Println(tipo(func() {}))
	fmt.Println(tipo(func() int { return 1 }))
}
