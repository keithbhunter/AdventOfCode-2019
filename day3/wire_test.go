package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewWire(t *testing.T) {
	w := newWire("R8,U5,L5,D3")
	expected := []coordinate{
		// R8
		coordinate{1, 0},
		coordinate{2, 0},
		coordinate{3, 0},
		coordinate{4, 0},
		coordinate{5, 0},
		coordinate{6, 0},
		coordinate{7, 0},
		coordinate{8, 0},
		// U5
		coordinate{8, 1},
		coordinate{8, 2},
		coordinate{8, 3},
		coordinate{8, 4},
		coordinate{8, 5},
		// L5
		coordinate{7, 5},
		coordinate{6, 5},
		coordinate{5, 5},
		coordinate{4, 5},
		coordinate{3, 5},
		// D3
		coordinate{3, 4},
		coordinate{3, 3},
		coordinate{3, 2},
	}
	assert.Equal(t, expected, w.coordinates)
}
