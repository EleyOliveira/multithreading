package main

import (
	"io"
	"log"
	"net/http"
	"time"
)

func main() {
	cepBusca := "06182030"

	canal := make(chan string)

	//time.Sleep(5 * time.Second)
	go func() {
		time.Sleep(3 * time.Second)
		req, err := http.Get("http://viacep.com.br/ws/" + cepBusca + "/json/")
		if err != nil {
			panic(err)
		}
		defer req.Body.Close()
		res, err := io.ReadAll(req.Body)
		if err != nil {
			panic(err)
		}
		canal <- string(res)
	}()

	go func() {
		time.Sleep(2 * time.Second)
		req, err := http.Get("https://brasilapi.com.br/api/cep/v1/" + cepBusca)
		if err != nil {
			panic(err)
		}
		defer req.Body.Close()
		res, err := io.ReadAll(req.Body)
		if err != nil {
			panic(err)
		}
		canal <- string(res)
	}()

	select {
	case resposta := <-canal:
		log.Println(resposta)
	case <-time.After(1 * time.Second):
		log.Println("timeout")
	}

}
