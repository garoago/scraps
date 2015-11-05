package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func fib(n int, results chan<- uint64) {
	defer close(results)
	a, b := uint64(0), uint64(1)
	results <- a
	for n > 0 {
		if b < a {
			panic(fmt.Sprintf("fib() overflow: a > b (%d > %d)", a, b))
		}
		a, b = b, a+b
		results <- a
		n--
	}
}

func parseArg() (n int) {
	var err error
	if len(os.Args) == 2 {
		n, err = strconv.Atoi(os.Args[1])
		if err == nil && n < 0 {
			err = errors.New("Argument must be n >= 0")
		}
	}
	if err != nil {
		fmt.Println(err)
	}
	if err != nil || len(os.Args) != 2 {
		fmt.Println("Usage:  fibochan <n>")
		os.Exit(1)
	}
	return n
}

func main() {
	n := parseArg()
	results := make(chan uint64)
	go fib(n, results)
	i := 0
	for r := range results {
		fmt.Printf("%3d: %d\n", i, r)
		i++
	}
}
