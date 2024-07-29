package main

import (
	"io"
	"log"
	"net/http"
	"time"
)

func main() {
	cepBusca := "06186110"

	canal := make(chan string, 2)

	go func() {
		req, err := http.Get("http://viacep.com.br/ws/" + cepBusca + "/json/")
		if err != nil {
			panic(err)
		}
		defer req.Body.Close()
		res, err := io.ReadAll(req.Body)
		if err != nil {
			panic(err)
		}
		canal <- "viacep"
		canal <- string(res)
	}()

	go func() {
		req, err := http.Get("https://brasilapi.com.br/api/cep/v1/" + cepBusca)
		if err != nil {
			panic(err)
		}
		defer req.Body.Close()
		res, err := io.ReadAll(req.Body)
		if err != nil {
			panic(err)
		}
		canal <- "brasilapi"
		canal <- string(res)
	}()

	select {
	case resposta := <-canal:
		log.Println(resposta)
		log.Println(<-canal)
	case <-time.After(1 * time.Second):
		log.Println("timeout")
	}

}
