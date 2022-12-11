package shape

import (
	"testing"

	shape "github.com/ViktorKram/Shape"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewTriangle_validSides_returnsTriangle(t *testing.T) {
	triangle, err := shape.NewTriangle(1, 2, 3)

	require.NotNil(t, triangle)
	require.NoError(t, err)
}

func TestNewTriangle_sideEqualsOrLessThanZero_returnsError(t *testing.T) {
	cases := []struct {
		sideA float64
		sideB float64
		sideC float64
	}{
		{sideA: 0,
			sideB: 1,
			sideC: 2},
		{sideA: 1,
			sideB: 0,
			sideC: 2},
		{sideA: 1,
			sideB: 2,
			sideC: 0},
	}

	for _, tCase := range cases {
		triangle, err := shape.NewTriangle(tCase.sideA, tCase.sideB, tCase.sideC)

		require.Nil(t, triangle)
		require.EqualError(t, err, "side is less or equals zero")
	}
}

func TestNewTriangle_invalidSidesSum_returnsError(t *testing.T) {
	triangle, err := shape.NewTriangle(1, 1, 10)

	require.Nil(t, triangle)
	require.EqualError(t, err, "triangle with these sides cannot exist")
}

func TestGetSideA_returnsSideA(t *testing.T) {
	triangle, _ := shape.NewTriangle(1, 2, 3)

	assert.Equal(t, float64(1), triangle.GetSideA())
}

func TestGetSideB_returnsSideB(t *testing.T) {
	triangle, _ := shape.NewTriangle(1, 2, 3)

	assert.Equal(t, float64(2), triangle.GetSideB())
}

func TestGetSideC_returnsSideC(t *testing.T) {
	triangle, _ := shape.NewTriangle(1, 2, 3)

	assert.Equal(t, float64(3), triangle.GetSideC())
}

func TestGetArea_returnCorrectArea(t *testing.T) {
	triangle, _ := shape.NewTriangle(3, 4, 5)

	assert.Equal(t, 2.449489742783178, triangle.GetArea())
}

func TestIsRightAngle_returnsTrue(t *testing.T) {
	triangle, _ := shape.NewTriangle(3, 4, 5)

	require.True(t, triangle.IsRightAngle())
}

func TestIsRightAngle_returnsFalse(t *testing.T) {
	triangle, _ := shape.NewTriangle(3, 4, 6)

	require.False(t, triangle.IsRightAngle())
}
