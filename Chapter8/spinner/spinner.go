package main

import (
	"fmt"
	"time"
)

func spinner(delay time.Duration) {
	for {
		for _, r := range `-\|/` {
			fmt.Printf("\r%c", r)
			time.Sleep(delay)
		}
	}
}

func fib(x int) int {
	if x < 2 {
		return x
	}
	return fib(x-2) + fib(x-1)
}

func main() {
	go spinner(100 * time.Millisecond)
	const n = 45
	fibN := fib(45)
	fmt.Printf("\rFibonacci(%d) = %d\n", n, fibN)
}
