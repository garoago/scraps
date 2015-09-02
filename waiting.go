package main

import (
	"fmt"
	"time"
)

func goroutine(done chan<- int) {
	fmt.Println("\tgorountine start")
	time.Sleep(2 * time.Second)
	fmt.Println("\tgorountine sending")
	done <- 1
	fmt.Println("\tgorountine end")
}

func main() {
	fmt.Println("main start")
	done := make(chan int)
	go goroutine(done)
	fmt.Println("main after `go`")
	//time.Sleep(2 * time.Second)
	<-done // blocks until `done` is written to
	fmt.Println("main received")
	//time.Sleep(2 * time.Second)
	fmt.Println("main end")
}
