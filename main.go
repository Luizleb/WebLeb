package main

import (
	"fmt"
	"net/http"
)

var db map[string]float32

func list(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("content-type", "text/html")
	for k, val := range db {
		fmt.Fprintf(w, "<h3>Item:  %s, valor : %v</h3>\n", k, val)
	}
}

func price(w http.ResponseWriter, r *http.Request) {
	reqItem := r.URL.Query().Get("item")
	price, ok := db[reqItem]
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "No such item: %q\n", reqItem)
		return
	}
	fmt.Fprintf(w, "Price for item : %s = %v\n", reqItem, price)
}

func main() {
	db = map[string]float32{
		"openGL": 100,
		"Golang": 65,
	}
	http.HandleFunc("/list", list)
	http.HandleFunc("/price", price)
	http.ListenAndServe(":3000", nil)
}
