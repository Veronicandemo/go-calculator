package constants

import "math"

var (
	e = &Constant{Name: "e",
		Value: math.E}
	pi    = &Constant{Name: "pi", Value: math.Pi}
	piSym = &Constant{Name: "π", Value: math.Pi}
	i     = &Constant{Name: "i", ComValue: complex(0, 1)}
)

func init() {
	Register(e)
	Register(pi)
	Register(piSym)
	Register(i)
}
