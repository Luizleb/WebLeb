package main

import (
	"fmt"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("content-type", "text/plain")
	fmt.Fprint(w, "<h2>Agora utilizando HandleFunc</h2>")
}

func teste(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("content-type", "text/html")
	fmt.Fprint(w, "<h2>Página de Teste</h2>")
}

func main() {
	http.HandleFunc("/home", home)
	http.HandleFunc("/teste/", teste)
	http.ListenAndServe(":3000", nil)
}
