package functions

import (
	"github.com/Veronicandemo/calculator/operators"
	"math"
)

var (
	sqrt = &operators.Operator{
		Name:          "sqrt",
		Precedence:    0,
		Associativity: operators.L,
		Args:          1,
		Operation: func(args []float64) float64 {
			return math.Sqrt(args[0])
		},
	}
)

func init() {
	Register(sqrt)
}
