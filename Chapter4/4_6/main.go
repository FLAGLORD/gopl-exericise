package main

import (
	"fmt"
	"unicode"
	"unicode/utf8"
)

func main() {
	s := []byte("a a aaa                我我")
	s = removeDuplicateSpace(s)
	fmt.Println(string(s))
}

func removeDuplicateSpace(str []byte) []byte {
	for i := 0; i < len(str); {
		unicodeChar, size := utf8.DecodeRune(str[i:])
		if unicode.IsSpace(unicodeChar) {
			pos := i + size
			following, followingSize := utf8.DecodeRune(str[pos:])
			for unicode.IsSpace(following) {
				pos += followingSize
				following, followingSize = utf8.DecodeRune(str[pos:])
			}
			copy(str[i+size:], str[pos:])
			// pos - (i + size) 表示跳过的空格所占的byte大小
			// len(str) - xxx 表示将跳过的部分删去
			str = str[:len(str)-(pos-(i+size))]
		}
		i += size
	}
	return str
}
