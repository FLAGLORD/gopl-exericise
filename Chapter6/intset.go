package main

import (
	"fmt"
	"strconv"
	"strings"
)

// An Intset is a set of small non-negative integers
// Its zero value represents the empty set
type IntSet struct {
	words []uint64
}

// Has checks whether the set contains the non-negative value x
func (s *IntSet) Has(x int) bool {
	word, bit := x/64, uint(x%64)
	return word <= len(s.words) && ((s.words[word] & (1 << bit)) != 0)
}

// Add add the non-negative value x to the set
func (s *IntSet) Add(x int) {
	word, bit := x/64, uint(x%64)
	// len(s.words) should at least be (word + Chapter1)
	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}
	s.words[word] |= 1 << bit
}

// UnionWith set s to the union of s and t
func (s *IntSet) UnionWith(t *IntSet) {
	for len(s.words) < len(t.words) {
		s.words = append(s.words, 0)
	}
	for i, tword := range t.words {
		s.words[i] |= tword
	}
}

// IntersectWith set s to the intersect of s and t
func (s *IntSet) IntersectWith(t *IntSet) {
	for len(s.words) < len(t.words) {
		s.words = append(s.words, 0)
	}
	for i, tword := range t.words {
		s.words[i] &= tword
	}
}

func (s *IntSet) DifferenceWith(t *IntSet) {
	// len(s) < len(t), only set len(s) elements
	// otherwise( len(s) >= len(t) ) only set len(t) elements,  对于多出来的部分必然是差集
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	minLen := min(len(s.words), len(t.words))
	for i := 0; i < minLen; i++ {
		intersect := s.words[i] & t.words[i]
		s.words[i] ^= intersect
	}
}

func (s *IntSet) SymmetricDifference(t *IntSet) {
	for len(s.words) < len(t.words) {
		s.words = append(s.words, 0)
	}
	for i, tword := range t.words {
		s.words[i] ^= tword
	}
}

// String returns the set as a string of the form "{Chapter1 2 3}"
func (s *IntSet) String() string {
	var builder strings.Builder
	builder.WriteByte('{')
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < 64; j++ {
			if word&(1<<j) != 0 {
				if builder.Len() > 1 {
					builder.WriteByte(' ')
				}
				builder.WriteString(strconv.Itoa(i*64 + j))
			}
		}
	}
	builder.WriteByte('}')
	return builder.String()
}

func (s *IntSet) Len() int {
	len := 0
	for _, word := range s.words {
		if word == 0 {
			continue
		}
		for word != 0 {
			len++
			word &= word - 1
		}
	}
	return len
}

// Remove x from the set s
func (s *IntSet) Remove(x int) {
	word, bit := x/64, uint64(x%64)
	// 元素一定不存在
	if word >= len(s.words) {
		return
	}
	// 首先 set 一下， x对应的位
	s.words[word] |= 1 << bit
	// 利用异或
	// 0 xor 0 = 0, 0 xor Chapter1 = Chapter1
	s.words[word] ^= 1 << bit
}

// Remove all elements  from the set
func (s *IntSet) Clear() {
	s.words = nil
}

// Return a copy of the set
func (s *IntSet) Copy() *IntSet {
	b := make([]uint64, len(s.words), len(s.words))
	copy(b, s.words)
	return &IntSet{
		words: b,
	}
}

// Add a group of elements
func (s *IntSet) AddAll(elements ...int) {
	for _, element := range elements {
		s.Add(element)
	}
}

func (s *IntSet) Elems() []uint64 {
	var elements []uint64
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < 64; j++ {
			if word&(1<<j) != 0 {
				elements = append(elements, uint64(i*64+j))
			}
		}
	}
	return elements
}

func main() {
	var x, y IntSet
	x.Add(1)
	x.Add(144)
	x.Add(9)

	y.Add(9)
	y.Add(42)

	fmt.Println(x.String())
	fmt.Println(x.Len())
	fmt.Println(y.String())
	fmt.Println(y.Len())

	x.UnionWith(&y)
	fmt.Println(x.String())
	fmt.Println(x.Len())
	x.DifferenceWith(&y)
	fmt.Println(x.String())
	for _, element := range x.Elems() {
		fmt.Println(element)
	}
}
