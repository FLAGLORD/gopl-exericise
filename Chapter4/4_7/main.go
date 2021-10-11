package main

import (
	"fmt"
	"unicode/utf8"
)

func reverseByte(b []byte) {
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}
}

func reverseUtf8(b []byte) {
	for i := 0; i < len(b); {
		_, size := utf8.DecodeRune(b[i:])
		reverseByte(b[i : i+size])
		i += size
	}
	reverseByte(b)
}

func main() {
	testStr := "永远滴神"
	reverseUtf8([]byte(testStr))
	fmt.Println(testStr)
	anotherTestStr := []byte("永远滴神")
	reverseUtf8(anotherTestStr)
	fmt.Println(string(anotherTestStr))
}
