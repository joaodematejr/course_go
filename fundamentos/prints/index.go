package main

import "fmt"

func main() {
	fmt.Print("Hello")
	fmt.Print("Hello")

	fmt.Println("Hello, World!")
	fmt.Println("Hello, World!")

	x := 3.141516

	//fmt.Printf("O valor de x é %.2f", x)

	xs := fmt.Sprint(x)
	fmt.Println("O valor de x é " + xs)
	fmt.Println("O valor de x é", x)
	fmt.Printf("O valor de x é %.2f", x)

	a := 1
	b := 1.9999
	c := false
	d := "opa"
	fmt.Printf("%d %f %v %s", a, b, c, d)
	fmt.Printf("%v %v %v %v", a, b, c, d)

}
