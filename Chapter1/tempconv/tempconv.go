package tempconv

import (
	"flag"
	"fmt"
)

type Celsius float64
type Fahrenheit float64
type Kelvin float64

//此处如果不显式声明类型，AbsoluteZero 根据值推断可能被认为是float64
const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC     Celsius = 0
	BoilingC      Celsius = 100
	AbsoluteZeroK Kelvin  = 0
)

func (c Celsius) String() string {
	return fmt.Sprintf("%g °C", c)
}

func (f Fahrenheit) String() string {
	return fmt.Sprintf("%g °F", f)
}

func (k Kelvin) String() string {
	return fmt.Sprintf("%g K", k)
}

type kelvinFlag struct {
	Kelvin
}

func (k *kelvinFlag) Set(s string) error {
	var unit string
	var value float64
	fmt.Sscanf(s, "%f%s", &value, &unit)
	switch unit {
	case "K":
		k.Kelvin = Kelvin(value)
		return nil
	}
	return fmt.Errorf("invalid temperature input: %s", s)
}

func KelvinFlag(name string, value Kelvin, usage string) *Kelvin {
	k := kelvinFlag{value}
	flag.CommandLine.Var(&k, name, usage)
	return &k.Kelvin
}
