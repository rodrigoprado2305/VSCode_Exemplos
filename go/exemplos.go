package main

import (
	"fmt"
	"runtime"
	"time"
)

func contador(tipo string) {
	for i := 0; i < 5; i++ {
		fmt.Println(tipo, i)
		time.Sleep(time.Second)
	}
}

func main() {
	runtime.GOMAXPROCS(1)
	fmt.Println("Comecou")

	go func() {
		for {

		}
	}()
	time.Sleep(time.Second)
	fmt.Println("Terminou")
}
