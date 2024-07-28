package main

import (
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", BuscaCepHandler)
	http.ListenAndServe(":8080", nil)

}

func BuscaCepHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	cepBusca := r.URL.Query().Get("cep")
	if cepBusca == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	req, err := http.Get("http://viacep.com.br/ws/" + cepBusca + "/json/")
	if err != nil {
		panic(err)
	}
	defer req.Body.Close()
	res, err := io.ReadAll(req.Body)
	if err != nil {
		panic(err)
	}
	w.Write(res)
	req, err = http.Get("https://brasilapi.com.br/api/cep/v1/" + cepBusca)
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
