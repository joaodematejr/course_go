package main

import "fmt"

func fatorial(n int) (int, error) {
	switch {
	case n < 0:
		return -1, fmt.Errorf("número inválido: %d", n)
	case n == 0:
		return 1, nil
	default:
		fatorialAnteriror, _ := fatorial(n - 1)
		return n * fatorialAnteriror, nil
	}
}

func main() {
	resultado, _ := fatorial(5)
	fmt.Println(resultado)
	_, err := fatorial(-4)
	if err != nil {
		fmt.Println(err)
	}

	// MA FORMA MELHOR SERIA ....

	// if resultado, err := fatorial(-4); err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	fmt.Println(resultado)
	// }

}
