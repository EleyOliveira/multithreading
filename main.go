package main

import (
	"io"
	"log"
	"net/http"
	"time"
)

func main() {
	cepBusca := "06182030"

	time.Sleep(10 * time.Second)
	func() {
		req, err := http.Get("http://viacep.com.br/ws/" + cepBusca + "/json/")
		if err != nil {
			panic(err)
		}
		defer req.Body.Close()
		res, err := io.ReadAll(req.Body)
		if err != nil {
			panic(err)
		}
		log.Println(string(res))
	}()

	func() {
		req, err := http.Get("https://brasilapi.com.br/api/cep/v1/" + cepBusca)
		if err != nil {
			panic(err)
		}
		defer req.Body.Close()
		res, err := io.ReadAll(req.Body)
		if err != nil {
			panic(err)
		}
		log.Println(string(res))
	}()
}
