package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMain(t *testing.T) {
	main()
}

// R8,U5,L5,D3
// U7,R6,D4,L4 = distance 6
func TestExample1(t *testing.T) {
	w1 := newWire("R8,U5,L5,D3")
	w2 := newWire("U7,R6,D4,L4")
	assert.Equal(t, 6, determineDistance(w1, w2))
}

// R75,D30,R83,U83,L12,D49,R71,U7,L72
// U62,R66,U55,R34,D71,R55,D58,R83    = distance 159
func TestExample2(t *testing.T) {
	w1 := newWire("R75,D30,R83,U83,L12,D49,R71,U7,L72")
	w2 := newWire("U62,R66,U55,R34,D71,R55,D58,R83")
	assert.Equal(t, 159, determineDistance(w1, w2))
}

// R98,U47,R26,D63,R33,U87,L62,D20,R33,U53,R51
// U98,R91,D20,R16,D67,R40,U7,R15,U6,R7        = distance 135
func TestExample3(t *testing.T) {
	w1 := newWire("R98,U47,R26,D63,R33,U87,L62,D20,R33,U53,R51")
	w2 := newWire("U98,R91,D20,R16,D67,R40,U7,R15,U6,R7")
	assert.Equal(t, 135, determineDistance(w1, w2))
}

func TestDistance(t *testing.T) {
	assert.Equal(t, 7, distance(coordinate{3, 4}))
}

// R8,U5,L5,D3
// U7,R6,D4,L4 = 30 steps
func TestStepsExample1(t *testing.T) {
	w1 := newWire("R8,U5,L5,D3")
	w2 := newWire("U7,R6,D4,L4")
	assert.Equal(t, 30, determineSteps(w1, w2))
}

// R75,D30,R83,U83,L12,D49,R71,U7,L72
// U62,R66,U55,R34,D71,R55,D58,R83    = 610 steps
func TestStepsExample2(t *testing.T) {
	w1 := newWire("R75,D30,R83,U83,L12,D49,R71,U7,L72")
	w2 := newWire("U62,R66,U55,R34,D71,R55,D58,R83")
	assert.Equal(t, 610, determineSteps(w1, w2))
}

// R98,U47,R26,D63,R33,U87,L62,D20,R33,U53,R51
// U98,R91,D20,R16,D67,R40,U7,R15,U6,R7        = 410 steps
func TestStepsExample3(t *testing.T) {
	w1 := newWire("R98,U47,R26,D63,R33,U87,L62,D20,R33,U53,R51")
	w2 := newWire("U98,R91,D20,R16,D67,R40,U7,R15,U6,R7")
	assert.Equal(t, 410, determineSteps(w1, w2))
}
