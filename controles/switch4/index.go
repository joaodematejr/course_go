package main

func tipo(i interface{}) string {
	switch i.(type) {
	case int:
		return "Inteiro"
	case float32, float64:
		return "Real"
	case string:
		return "String"
	case func():
		return "Função"
	default:
		return "Tipo não identificado"
	}

}

func main() {
	println(tipo(9.8))
	println(tipo(6))
	println(tipo("Opa"))
	println(tipo(func() {}))
	println(tipo(map[string]int{}))
}
