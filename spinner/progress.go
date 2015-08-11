// progress.go

package main

import (
	"fmt"
	"time"
)

func progress(done <-chan bool) {
	for {
		select {
		case <-done:
			break
		default:
			fmt.Print(".")
			time.Sleep(100 * time.Millisecond)
		}
	}
}

func slowFunction() int {
	time.Sleep(3 * time.Second)
	return 42
}

func main() {
	done := make(chan bool)
	go progress(done)
	result := slowFunction()
	done <- true
	fmt.Printf("Answer %d\n", result)
}
