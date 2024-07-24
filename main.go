package main

import "net/http"

func main() {
	http.HandleFunc("/", cepHandler)
	http.ListenAndServe(":8080", nil)

}

func cepHandler(w http.ResponseWriter, r *http.Request) {

}
