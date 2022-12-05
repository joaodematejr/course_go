package main

import "fmt"

func main() {
	ch := make(chan int, 1)

	ch <- 1 // envia um valor para o canal (escrita)

	<-ch // recebe um valor do canal (leitura)

	ch <- 2 // envia um valor para o canal (escrita)

	fmt.Println(<-ch) // recebe um valor do canal (leitura) e imprime na tela

}
