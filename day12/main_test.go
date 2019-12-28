package main

import (
	"io/ioutil"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestPart1(t *testing.T) {
	bytes, err := ioutil.ReadFile("input.txt")
	require.NoError(t, err)

	str := string(bytes)
	moons := newMoons(str)
	energy := totalEnergyAfter(1000, moons)
	assert.Equal(t, 7636, energy)
}

func TestNewMoons(t *testing.T) {
	input := `<x=-1, y=0, z=2>
	<x=2, y=-10, z=-7>
	<x=4, y=-8, z=8>
	<x=3, y=5, z=-1>`
	input = strings.ReplaceAll(input, "\t", "")
	moons := newMoons(input)

	assert.Len(t, moons, 4)
	assert.Equal(t, point{-1, 0, 2}, moons[0].position)
	assert.Equal(t, point{2, -10, -7}, moons[1].position)
	assert.Equal(t, point{4, -8, 8}, moons[2].position)
	assert.Equal(t, point{3, 5, -1}, moons[3].position)
}

func TestSimulateMotion(t *testing.T) {
	input := `<x=-1, y=0, z=2>
	<x=2, y=-10, z=-7>
	<x=4, y=-8, z=8>
	<x=3, y=5, z=-1>`
	input = strings.ReplaceAll(input, "\t", "")
	moons := newMoons(input)

	t.Run("1 iteration", func(t *testing.T) {
		result := simulateMotion(moons, 1)

		assert.Equal(t, 2, result[0].position.x)
		assert.Equal(t, -1, result[0].position.y)
		assert.Equal(t, 1, result[0].position.z)
		assert.Equal(t, 3, result[0].velocity.x)
		assert.Equal(t, -1, result[0].velocity.y)
		assert.Equal(t, -1, result[0].velocity.z)

		assert.Equal(t, 3, result[1].position.x)
		assert.Equal(t, -7, result[1].position.y)
		assert.Equal(t, -4, result[1].position.z)
		assert.Equal(t, 1, result[1].velocity.x)
		assert.Equal(t, 3, result[1].velocity.y)
		assert.Equal(t, 3, result[1].velocity.z)

		assert.Equal(t, 1, result[2].position.x)
		assert.Equal(t, -7, result[2].position.y)
		assert.Equal(t, 5, result[2].position.z)
		assert.Equal(t, -3, result[2].velocity.x)
		assert.Equal(t, 1, result[2].velocity.y)
		assert.Equal(t, -3, result[2].velocity.z)

		assert.Equal(t, 2, result[3].position.x)
		assert.Equal(t, 2, result[3].position.y)
		assert.Equal(t, 0, result[3].position.z)
		assert.Equal(t, -1, result[3].velocity.x)
		assert.Equal(t, -3, result[3].velocity.y)
		assert.Equal(t, 1, result[3].velocity.z)
	})

	t.Run("2 iterations", func(t *testing.T) {
		result := simulateMotion(moons, 2)

		assert.Equal(t, 5, result[0].position.x)
		assert.Equal(t, -3, result[0].position.y)
		assert.Equal(t, -1, result[0].position.z)
		assert.Equal(t, 3, result[0].velocity.x)
		assert.Equal(t, -2, result[0].velocity.y)
		assert.Equal(t, -2, result[0].velocity.z)

		assert.Equal(t, 1, result[1].position.x)
		assert.Equal(t, -2, result[1].position.y)
		assert.Equal(t, 2, result[1].position.z)
		assert.Equal(t, -2, result[1].velocity.x)
		assert.Equal(t, 5, result[1].velocity.y)
		assert.Equal(t, 6, result[1].velocity.z)

		assert.Equal(t, 1, result[2].position.x)
		assert.Equal(t, -4, result[2].position.y)
		assert.Equal(t, -1, result[2].position.z)
		assert.Equal(t, 0, result[2].velocity.x)
		assert.Equal(t, 3, result[2].velocity.y)
		assert.Equal(t, -6, result[2].velocity.z)

		assert.Equal(t, 1, result[3].position.x)
		assert.Equal(t, -4, result[3].position.y)
		assert.Equal(t, 2, result[3].position.z)
		assert.Equal(t, -1, result[3].velocity.x)
		assert.Equal(t, -6, result[3].velocity.y)
		assert.Equal(t, 2, result[3].velocity.z)
	})

	t.Run("10 iterations", func(t *testing.T) {
		result := simulateMotion(moons, 10)

		assert.Equal(t, 2, result[0].position.x)
		assert.Equal(t, 1, result[0].position.y)
		assert.Equal(t, -3, result[0].position.z)
		assert.Equal(t, -3, result[0].velocity.x)
		assert.Equal(t, -2, result[0].velocity.y)
		assert.Equal(t, 1, result[0].velocity.z)

		assert.Equal(t, 1, result[1].position.x)
		assert.Equal(t, -8, result[1].position.y)
		assert.Equal(t, 0, result[1].position.z)
		assert.Equal(t, -1, result[1].velocity.x)
		assert.Equal(t, 1, result[1].velocity.y)
		assert.Equal(t, 3, result[1].velocity.z)

		assert.Equal(t, 3, result[2].position.x)
		assert.Equal(t, -6, result[2].position.y)
		assert.Equal(t, 1, result[2].position.z)
		assert.Equal(t, 3, result[2].velocity.x)
		assert.Equal(t, 2, result[2].velocity.y)
		assert.Equal(t, -3, result[2].velocity.z)

		assert.Equal(t, 2, result[3].position.x)
		assert.Equal(t, 0, result[3].position.y)
		assert.Equal(t, 4, result[3].position.z)
		assert.Equal(t, 1, result[3].velocity.x)
		assert.Equal(t, -1, result[3].velocity.y)
		assert.Equal(t, -1, result[3].velocity.z)
	})
}

func TestTotalEnergyAfter(t *testing.T) {
	input := `<x=-1, y=0, z=2>
	<x=2, y=-10, z=-7>
	<x=4, y=-8, z=8>
	<x=3, y=5, z=-1>`
	input = strings.ReplaceAll(input, "\t", "")
	moons := newMoons(input)
	energy := totalEnergyAfter(10, moons)
	assert.Equal(t, 179, energy)
}
