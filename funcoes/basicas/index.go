package main

import "fmt"

func funcao1() {
	fmt.Println("F1")
}

func funcao2(p1, p2 string) {
	fmt.Printf("F2: %s %s", p1, p2)
}

func funcao3() string {
	return "F3"
}

func funcao4(p1, p2 string) string {
	return fmt.Sprintf("F4: %s %s", p1, p2)
}

func funcao5() (string, string) {
	return "Retorno 1", "Retorno 2"
}

func main() {
	funcao1()
	funcao2("Param1", "Param2")
	r3, r4 := funcao3(), funcao4("Param1", "Param2")
	fmt.Println(r3)
	fmt.Println(r4)
	r51, r52 := funcao5()
	fmt.Println(r51, r52)

}
