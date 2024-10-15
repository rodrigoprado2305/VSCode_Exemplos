package main

import (
	"fmt"
	"time"
)

func main() {
	queue := make(chan int)

	go func() { //compartilhando o channel
		i := 0
		for {
			time.Sleep(time.Second)
			queue <- i
			i++
		}
	}()

	for x := range queue {
		fmt.Println(x)
	}
}
