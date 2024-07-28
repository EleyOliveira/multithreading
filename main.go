package main

import (
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", cepHandler)
	http.ListenAndServe(":8080", nil)

}

func cepHandler(w http.ResponseWriter, r *http.Request) {
	req, err := http.Get("http://viacep.com.br/ws/06182030/json/")
	if err != nil {
		panic(err)
	}
	defer req.Body.Close()
	res, err := io.ReadAll(req.Body)
	if err != nil {
		panic(err)
	}
	w.Write(res)
	req, err = http.Get("https://brasilapi.com.br/api/cep/v1/62051032")
	if err != nil {
		panic(err)
	}
	defer req.Body.Close()
	res, err = io.ReadAll(req.Body)
	if err != nil {
		panic(err)
	}
	w.Write(res)
}
