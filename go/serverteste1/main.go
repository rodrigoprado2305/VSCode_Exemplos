package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/", teste)

	err := http.ListenAndServe(":3001", nil)

	if err != nil {

		log.Fatal("ListenAndServe: ", err)

	}

}

func teste(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1> %s %d", "  <h1> Teste </h1>", 1070)
}

func cliente(w http.ResponseWriter, r *http.Request) {
	//http.HandleFunc("/", "cliente.html")
}
