package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
)

func encaminhar(origem <-chan string, destino chan<- string) {
	for {
		destino <- <-origem
	}
}

// Multiplexar - Multiplexar é a técnica de combinar várias fontes de dados em uma única fonte de dados.

func juntar(entrada1, entrada2 <-chan string) <-chan string {
	c := make(chan string)
	go encaminhar(entrada1, c)
	go encaminhar(entrada2, c)
	return c
}

// <-chan - canal somente-leitura
func Titulo(urls ...string) <-chan string {
	c := make(chan string)

	for _, url := range urls {
		go func(cUrl string) {
			resp, _ := http.Get(cUrl)
			html, _ := ioutil.ReadAll(resp.Body)

			r, _ := regexp.Compile("<title>(.*?)<\\/title>")
			aRetorno := r.FindStringSubmatch(string(html))

			// teste para evitar erro
			if cap(aRetorno) == 0 {
				c <- "Erro ao ler página " + cUrl
				return
			}

			c <- aRetorno[1]
		}(url)
	}

	return c
}

func main() {
	c := juntar(
		Titulo("https://www.cod3r.com.br", "https://www.google.com"),
		Titulo("https://www.amazon.com", "https://www.youtube.com"),
	)
	fmt.Println("Primeiros:", <-c, "|", <-c)
	fmt.Println("Segundos:", <-c, "|", <-c)
}
