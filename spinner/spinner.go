// spinner.go

package main

import (
	"fmt"
	"strings"
	"time"
)

func spin(msg string, done <-chan bool) {
	const spinChars = "|/-\\"
	backspaces := strings.Repeat("\x08", len(msg)+2)

Forever:
	for {
		for _, char := range spinChars {
			select {
			case <-done:
				break Forever
			default:
				fmt.Printf("%c %s%s", char, msg, backspaces)
				time.Sleep(100 * time.Millisecond)
			}
		}
	}
	// clear spin line -- not working at all...
	fmt.Printf("%s%s", strings.Repeat(" ", len(msg)+2), backspaces)
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
	fmt.Printf("Answer %d\n", result)
}
