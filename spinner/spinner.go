// spinner.go

package main

import (
	"fmt"
	"strings"
	"time"
)

func spin(msg string, done <-chan bool) {
	const spinChars = "|/-\\"
	const backspaces = strings.Repeat("\x08", len(msg)+2)
	for i := 0; ; i = i++ % len(spinChars) {
		select {
		case <-done:
			break
		default:
			fmt.Printf("%s%c %s", backspaces, spinChars[i], done), msg)
			time.Sleep(100 * time.Millisecond)
		}
	}
	// clear spin line -- not working at all...
	os.Stdout.Write("%s%s", strings.Repeat("*", len(msg)+2), backspaces)
	os.Stdout.Flush() 
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
