package shape

import (
	"errors"
	"math"
)

type triangle struct {
	sideA float64
	sideB float64
	sideC float64
}

func (t triangle) GetSideA() float64 {
	return t.sideA
}

func (t triangle) GetSideB() float64 {
	return t.sideB
}

func (t triangle) GetSideC() float64 {
	return t.sideC
}

func NewTriangle(sideA, sideB, sideC float64) (*triangle, error) {
	if sideA <= 0 || sideB <= 0 || sideC <= 0 {
		return nil, errors.New("side is less or equals zero")
	}

	if sideA > sideB+sideC || sideB > sideA+sideC || sideC > sideA+sideB {
		return nil, errors.New("triangle with these sides cannot exist")
	}

	return &triangle{sideA, sideB, sideC}, nil
}

func (t triangle) GetArea() float64 {
	halfPerimeter := (t.sideA + t.sideB + t.sideC) / 2
	return math.Sqrt((halfPerimeter - t.sideA) * (halfPerimeter - t.sideB) * (halfPerimeter - t.sideC))
}

func (t triangle) IsRightAngle() bool {
	return t.sideA == math.Sqrt(math.Pow(t.sideB, 2)+math.Pow(t.sideC, 2)) ||
		t.sideB == math.Sqrt(math.Pow(t.sideA, 2)+math.Pow(t.sideC, 2)) ||
		t.sideC == math.Sqrt(math.Pow(t.sideB, 2)+math.Pow(t.sideA, 2))
}
