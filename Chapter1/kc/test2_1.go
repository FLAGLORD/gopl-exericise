package main

import (
	"exercise/Chapter1/tempconv"
	"flag"
)
import "fmt"

func main() {
	var k = tempconv.KelvinFlag("temp", 20.0, "the temperature")
	//for _, arg := range os.Args[Chapter1:] {
	//	t, err := strconv.ParseFloat(arg, 64)
	//	if err != nil {
	//		fmt.Fprintf(os.Stderr, "kc: %v\n", err)
	//		os.Exit(Chapter1)
	//	}
	//	k := tempconv.Kelvin(t)
	//	c := tempconv.Celsius(t)
	//
	//	fmt.Printf("%s = %s, %s = %s\n",
	//		k, tempconv.KToC(k), c, tempconv.CToK(c))
	//}
	flag.Parse()
	fmt.Println(*k)
}
