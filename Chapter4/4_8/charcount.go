package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
)

type class string

const (
	Letter class = "Letter"
	Digit  class = "Digit"
	Others class = "others"
)

func main() {
	statistics := make(map[class]int)

	input := bufio.NewReader(os.Stdin)
	for {
		r, _, err := input.ReadRune()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "charCocunt: %v\n", err)
			os.Exit(1)
		}
		if unicode.IsLetter(r) {
			statistics[Letter]++
		} else if unicode.IsDigit(r) {
			statistics[Digit]++
		} else {
			statistics[Others]++
		}
	}
	fmt.Println(statistics)
}
