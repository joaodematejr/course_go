package main

import "fmt"

func main() {
	p1 := Ponto{2.0, 2, 0}
	p2 := Ponto{2.0, 10.0}

	fmt.Println(catetos(p1, p2))
	fmt.Println(Distance(p1, p2))
}
