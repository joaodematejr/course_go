package main

func main() {
	http.HandleFunc("/usuarios/", UsuarioHandler)
	log.Println("Executando...")
	log.Fatal(http.ListenAndServe(":3000", nil))
}
