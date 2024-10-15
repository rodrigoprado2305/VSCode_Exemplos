package main

import (
	"fmt"
	"net/http"
)

func main() {
	res, err := http.Get("www.google.com")

	if err != nil {
		panic("Erro")
	}

	fmt.Println(res.Body)
}
