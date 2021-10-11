package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	input := bufio.NewScanner(os.Stdin)
	input.Split(bufio.ScanWords)
	freqTable := make(map[string]int)
	for input.Scan() {
		// Text() method returns string
		word := input.Text()
		freqTable[strings.ToLower(word)]++
	}
	for word, freq := range freqTable {
		fmt.Printf("%s : %d\n", word, freq)
	}
}
