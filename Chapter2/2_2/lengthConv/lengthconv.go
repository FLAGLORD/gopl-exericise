package lengthConv

import "fmt"

type Meter float64
type Foot float64

func (m Meter) String() string {
	return fmt.Sprintf("%g m", m)
}

func (f Foot) String() string {
	return fmt.Sprintf("%g ft", f)
}

func MToF(m Meter) Foot {
	return Foot(3.2808 * m)
}

func FToM(ft Foot) Meter {
	return Meter(ft / 3.2808)
}
