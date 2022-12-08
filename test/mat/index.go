package mat

import (
	"fmt"
	"strconv"
)

// MÉDIA É RESPONSÁVEL POR CALCULAR O QUE VOCE QUISER
func Media(numeros ...float64) float64 {
	total := 0.0
	for _, numero := range numeros {
		total += numero
	}
	media := total / float64(len(numeros))
	mediaArrondada, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", media), 64)
	return mediaArrondada
}
