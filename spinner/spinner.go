// spinner.go

package main

import (
	"fmt"
	"strings"
	"time"
)

func spin(msg string, done <-chan bool) {
	for {
		for _, char := range "|/-\\" {
			select {
			case <-done:
				return
			default:
				backspaces := strings.Repeat("\x08", len(msg)+2)
				fmt.Printf("%s%c %s", backspaces, char, msg)
				time.Sleep(100 * time.Millisecond)
			}
		}
	}
}

func slowFunction() int {
	time.Sleep(3 * time.Second)
	return 42
}

func main() {
	done := make(chan bool)
	go spin("thinking!", done)
	result := slowFunction()
	done <- true
	fmt.Printf("\nAnswer %d\n", result)
}
