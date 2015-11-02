/*
The incredibly inefficient O(2**n) recursive Fibonacci number generator
*/

package main

import (
	"fmt"
	"os"
	"strconv"
)

var callCount int

func fib(n int) int {
	callCount++
	// fmt.Printf("→%d ", n)  // uncomment to see every call!
	switch n {
	case 0:
		return 0
	case 1:
		return 1
	default:
		return fib(n-1) + fib(n-2)
	}
}

func ordinalSuffix(n int) string {
	if n >= 11 && n <= 19 {
		return "th"
	} else {
		switch n % 10 {
		case 1:
			return "st"
		case 2:
			return "nd"
		case 3:
			return "rd"
		default:
			return "th"
		}
	}
}

func main() {
	n := -1
	if len(os.Args) == 2 {
		var err error
		n, err = strconv.Atoi(os.Args[1])
		if err != nil {
			n = -1
			fmt.Println(err)
		}
	}
	if n == -1 {
		fmt.Println("Usage:  fiborec <n>")
		os.Exit(1)
	}

	fmt.Printf("\n%d%s number of the Fibonacci sequence: %d\n", n, ordinalSuffix(n), fib(n))
	fmt.Printf("(needed %d calls to fib() to compute)\n", callCount)
}
