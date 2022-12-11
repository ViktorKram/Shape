package shape

import (
	"math"
	"testing"

	shape "github.com/ViktorKram/Shape"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewCircle_validRadius_returnsCircle(t *testing.T) {
	circle, err := shape.NewCircle(1)

	require.NotNil(t, circle)
	require.NoError(t, err)
}

func TestNewCircle_radiusEqualsOrLessThanZero_returnsError(t *testing.T) {
	cases := []struct {
		radius float64
	}{
		{radius: 0},
		{radius: -1},
	}

	for _, tCase := range cases {
		circle, err := shape.NewCircle(tCase.radius)

		require.Nil(t, circle)
		require.EqualError(t, err, "radius is less or equals zero")
	}
}

func TestGetRadius_returnCorrectRadius(t *testing.T) {
	circle, _ := shape.NewCircle(1)

	assert.Equal(t, float64(1), circle.GetRadius())
}

func TestGetArea_validRadius_returnCorrectArea(t *testing.T) {
	circle, _ := shape.NewCircle(1)

	assert.Equal(t, math.Pi, circle.GetArea())
}
