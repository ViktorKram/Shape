package shape

import (
	"errors"
	"math"
)

type circle struct {
	radius float64
}

func (c circle) GetRadius() float64 {
	return c.radius
}

func NewCircle(radius float64) (*circle, error) {
	if radius <= 0 {
		return nil, errors.New("radius is less or equals zero.")
	}

	return &circle{radius}, nil
}

func (c circle) GetArea() float64 {
	return math.Pow(c.radius, 2) * math.Pi
}
