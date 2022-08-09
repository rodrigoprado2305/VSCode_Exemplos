package main

import (
	"fmt"
	"time"
)

func contador(tipo string) {
	for i := 0; i < 5; i++ {
		fmt.Println(tipo, i)
		time.Sleep(time.Second)
	}
}

func main() {
	//	contador("a")
	//  contador("b")

	go contador("c")
	go contador("d")
	time.Sleep(time.Second * 10)
}
