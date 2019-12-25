package shape

import (
	"math"
)

type Circle struct {
	Radius float64
}

// type Shape interface {
// 	Area() float64
// }

func (c *Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}
