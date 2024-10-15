package main

import (
	"fmt"
	"time"
)

func worker(workerid int, msg chan int) {
	for res := range msg {
		fmt.Println("worker: ", workerid, " Msg:", res)
		time.Sleep(time.Second)
	}
}

func main() {
	msg := make(chan int)
	go worker(1, msg) // no apache o maximo de worker são 3
	go worker(2, msg)
	go worker(3, msg)
	go worker(4, msg)
	go worker(5, msg)
	go worker(7, msg) //pode ter muitos workers

	for i := 0; i < 20; i++ {
		msg <- i
	}
}
