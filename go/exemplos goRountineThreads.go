package main

import (
	"fmt"
	"time"
)

// func nome_e_idade(nome string, idade int) (string, int, error) {
func contador_teste(tipo string, tempo int) {
	for i := 0; i < tempo; i++ {
		fmt.Println(tipo, i)
	}
	time.Sleep(time.Second * time.Duration(tempo))
}

func main() {
	fmt.Println("a")
	go contador_teste("b", 3)
	fmt.Println("c")
	contador_teste("d", 4)
	time.Sleep(time.Second)
}
