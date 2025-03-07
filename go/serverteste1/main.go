package main

import (
	"fmt"
	"log"
	"net/http"
)

//Para testar digitar no browser
//http://localhost:3001/

func main() {

	http.HandleFunc("/", teste)

	err := http.ListenAndServe(":3001", nil)

	if err != nil {

		log.Fatal("ListenAndServe: ", err)

	}

}

func teste(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprintf(w, "<h1> %s %d", "  <h1> Teste </h1>", 1070)
	fmt.Fprintf(w, "<h1>Teste %d</h1>", 1070)
}

//func abrir(w http.ResponseWriter, r *http.Request) {
//ttp.HandleFunc("/", "cliente.html")
//}
