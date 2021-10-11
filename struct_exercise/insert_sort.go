package main

import "fmt"

type tree struct {
	value       int
	left, right *tree
}

func sort(values []int) []int {
	var root *tree
	for _, val := range values {
		root = add(root, val)
	}

	newArr := make([]int, 0, len(values))
	newArr = appendValues(newArr, root)
	appendValues(values[:0], root)
	fmt.Println(newArr)
	return newArr
}
func appendValues(values []int, root *tree) []int {
	if root == nil {
		return values
	}
	values = appendValues(values, root.left)
	values = append(values, root.value)
	values = appendValues(values, root.right)
	return values

}

func add(t *tree, value int) *tree {
	if t == nil {
		return &tree{
			value: value,
		}
	}
	if value < t.value {
		t.left = add(t.left, value)
	} else {
		t.right = add(t.right, value)
	}
	return t
}

func main() {
	arr := []int{7, 7, 2, 1, 9, 4, 10}
	fmt.Println(sort(arr))
	fmt.Println(arr)
}
