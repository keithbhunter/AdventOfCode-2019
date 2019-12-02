package fuelcounterupper

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestExamples(t *testing.T) {
	// For a mass of 12, divide by 3 and round down to get 4, then subtract 2 to get 2.
	t.Run("mass of 12", func(t *testing.T) {
		assert.Equal(t, FuelRequired(12), 2)
	})
	
	// For a mass of 14, dividing by 3 and rounding down still yields 4, so the fuel required is also 2.
	t.Run("mass of 12", func(t *testing.T) {
		assert.Equal(t, FuelRequired(14), 2)
	})
	
	// For a mass of 1969, the fuel required is 654.
	t.Run("mass of 12", func(t *testing.T) {
		assert.Equal(t, FuelRequired(1969), 654)
	})
	
	// For a mass of 100756, the fuel required is 33583.
	t.Run("mass of 12", func(t *testing.T) {
		assert.Equal(t, FuelRequired(100756), 33583)
	})
}
