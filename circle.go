package shape

import "math"

type circle struct {
	radius float64
}

func (c circle) NewCircle(radius float64) circle {
	if radius <= 0 {
		panic("radius is less or equals zero")
	}

	return circle{radius}
}

func (c circle) GetArea() float64 {
	return math.Pow(c.radius, 2) * math.Pi
}
