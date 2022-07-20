package shape

import "math"

type triangle struct {
	sideA float64
	sideB float64
	sideC float64
}

func (t triangle) NewTriangle(sideA, sideB, sideC float64) triangle {
	if sideA <= 0 || sideB <= 0 || sideC <= 0 {
		panic("side is less or equals zero.")
	}

	if sideA > sideB+sideC || sideB > sideA+sideC || sideC > sideA+sideB {
		panic("triangle with these sides cannot exist")
	}

	return triangle{sideA, sideB, sideC}
}

func (t triangle) GetArea() float64 {
	halfPerimeter := (t.sideA + t.sideB + t.sideC) / 2
	return math.Sqrt((halfPerimeter - t.sideA) * (halfPerimeter - t.sideB) * (halfPerimeter - t.sideC))
}
