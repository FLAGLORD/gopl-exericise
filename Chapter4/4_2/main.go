package main

import (
	"bufio"
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
	"os"
)

func main() {
	SHAFlag := flag.String("shaxxx", "sha256", "shaxxx = sha256 | sha512 | sha384")
	flag.Parse()
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		inputContent := input.Bytes()
		switch *SHAFlag {
		case "sha256":
			fmt.Printf("sha256: %x \n", sha256.Sum256(inputContent))
		case "sha384":
			fmt.Printf("sha384: %x \n", sha512.Sum384(inputContent))
		case "sha512":
			fmt.Printf("sha512: %x \n", sha512.Sum512(inputContent))
		}
	}
}
