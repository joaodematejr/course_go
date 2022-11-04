package main

import "fmt"

func main() {
	funcsPorLetra := map[string]map[string]float64{
		"G": {
			"Guga":    11325.45,
			"Gabi":    15456.78,
			"Gustavo": 1200.00,
		},
		"J": {
			"José":   11325.45,
			"João":   15456.78,
			"Junior": 1200.00,
		},
		"P": {
			"Pedro": 11325.45,
			"Paulo": 15456.78,
			"Paula": 1200.00,
		},
	}

	delete(funcsPorLetra, "P")

	for letra, funcs := range funcsPorLetra {
		fmt.Println(letra, funcs)
	}
}
