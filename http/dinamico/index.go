package main

import (
	"fmt"
	"net/http"
	"time"
)

func horaCerta(w http.ResponseWriter, r *http.Request) {
	s := time.Now().Format("02/01/2006 15:04:05")
	fmt.Fprintf(w, "<h1>Hora Certa: %s</h1>", s)
}

func main() {
	http.HandleFunc("/hora-certa", horaCerta)
	fmt.Println("Executando...")
	http.ListenAndServe(":3000", nil)

}
