package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"time"
)

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

func oMaisRapido(url1, url2, url3 string) string {
	c1 := Titulo(url1)
	c2 := Titulo(url2)
	c3 := Titulo(url3)

	select {
	case t1 := <-c1:
		return t1
	case t2 := <-c2:
		return t2
	case t3 := <-c3:
		return t3
	case <-time.After(1000 * time.Millisecond):
		return "Todos perderam!"
		/* 	default:
		return "Sem resposta ainda!" */
	}
}

func main() {
	campeao := oMaisRapido(
		"https://www.cod3r.com.br",
		"https://www.google.com",
		"https://www.amazon.com",
	)
	fmt.Println(campeao)
}
