package main

import (
	"exercise/Chapter2/2_2/lengthConv"
	"fmt"
	"strconv"
)

import "os"

func main() {
	for _, arg := range os.Args[1:] {
		l, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "lengthConv: %v\n", err)
		}
		m := lengthConv.Meter(l)
		ft := lengthConv.Foot(l)

		fmt.Printf("%s = %s, %s = %s\n",
			m, lengthConv.MToF(m), ft, lengthConv.FToM(ft),
		)
	}

}
