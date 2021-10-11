package main

import "fmt"

func removeDuplicate(arr []int) []int {
	var i int
	for j := 0; j < len(arr); j++ {
		if j != 0 && arr[j] == arr[j-1] {
			j++
		}
		arr[i] = arr[j]
		i++
	}
	return arr[:i]
}

func main() {
	s := []int{1, 1, 2, 3, 3, 4}
	s = removeDuplicate(s)
	fmt.Println(s)
}
