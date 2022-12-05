package main

import (
	"fmt"
	"time"
)

func fale(pessoa, texto string, qtde int) {
	for i := 0; i < qtde; i++ {
		time.Sleep(time.Second)
		fmt.Printf("%s: %s (iteração %d)\r", pessoa, texto, i+1)
	}
}

func main() {
	//fale("Maria", "Pq vc não fala comigo?", 3)
	//fale("João", "Só posso falar depois de vc!", 1)

	//go fale("Maria", "Ei...", 10)
	//go fale("João", "Opa...!", 10)
	//time.Sleep(time.Second + 5)
	go fale("João", "Opa...!", 10)
	fale("Maria", "Ei...", 10)
	fmt.Println("Fim!")
}
