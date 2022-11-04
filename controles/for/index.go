package main

import (
	"fmt"
	"time"
)

func main() {
	i := 1
	for i <= 10 { // não tem while em go
		fmt.Println(i)
		i++
	}

	for i := 1; i <= 10; i++ { // for tradicional
		fmt.Println(i)
	}

	fmt.Println("Misturando...")

	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Println("Par")
		} else {
			fmt.Println("Ímpar")
		}
	}

	for {
		fmt.Println("Para sempre...")
		time.Sleep(time.Second * 3)
	}

}
