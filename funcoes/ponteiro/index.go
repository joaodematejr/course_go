package main

// revisao: um ponteiro e representado por um *

func inc1(n int) {
	n++
}

func inc2(n *int) {
	// revisao: * e usado para desreferenciar, ou seja,
	// ter acesso ao valor no qual o ponteiro aponta
	*n++
}

func main() {
	n := 1
	inc1(n) // por valor
	println(n)

	// revisao: & usado para obter o endereco da variavel
	println(&n)
	inc2(&n) // por referencia

	println(n)

}
