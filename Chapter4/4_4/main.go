package main

import "fmt"

func main() {
	s := []int{1, 2, 3, 4, 5}
	rotate(s, 2)
	fmt.Println(s)
}

func rotate(arr []int, n int) {
	n %= len(arr)
	arrLen := len(arr)
	tmp := append(arr, arr[:arrLen-n]...)
	copy(arr, tmp[arrLen-n:])
}
