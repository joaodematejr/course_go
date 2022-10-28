package main

import (
	"math/rand"
	"time"
)

func numeroAleatorio() int {
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)
	return r.Intn(10)
}

func main() {
	if i := numeroAleatorio(); i > 5 { // tb pode ser usado no switch
		println("Ganhou!")
	} else {
		println("Perdeu!")
	}

}
