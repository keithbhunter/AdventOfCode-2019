package main

import (
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestPart1And2(t *testing.T) {
	f, err := os.Open("input.txt")
	require.NoError(t, err)

	locations, err := getAsteroidLocations(f)
	require.NoError(t, err)

	asteroid, count := findBestLocation(locations)
	assert.Equal(t, point{26, 29}, asteroid)
	assert.Equal(t, 299, count)

	p := vaporizeAsteroids(asteroid, locations)
	assert.Equal(t, point{14, 19}, p[199])
}

func TestBestLocations(t *testing.T) {
	t.Run("example 1", func(t *testing.T) {
		in := `.#..#
		.....
		#####
		....#
		...##`
		in = strings.ReplaceAll(in, "\t", "")
		r := strings.NewReader(in)

		locations, err := getAsteroidLocations(r)
		require.NoError(t, err)

		asteroid, count := findBestLocation(locations)
		assert.Equal(t, point{3, 4}, asteroid)
		assert.Equal(t, 8, count)
	})

	t.Run("example 2", func(t *testing.T) {
		in := `......#.#.
		#..#.#....
		..#######.
		.#.#.###..
		.#..#.....
		..#....#.#
		#..#....#.
		.##.#..###
		##...#..#.
		.#....####`
		in = strings.ReplaceAll(in, "\t", "")
		r := strings.NewReader(in)

		locations, err := getAsteroidLocations(r)
		require.NoError(t, err)

		asteroid, count := findBestLocation(locations)
		assert.Equal(t, point{5, 8}, asteroid)
		assert.Equal(t, 33, count)
	})
}

func TestGetAsteroidLocations(t *testing.T) {
	in := `..#.#
	.....
	#####
	....#
	...##`
	in = strings.ReplaceAll(in, "\t", "")
	r := strings.NewReader(in)

	expected := []point{
		{2, 0},
		{4, 0},
		{0, 2},
		{1, 2},
		{2, 2},
		{3, 2},
		{4, 2},
		{4, 3},
		{3, 4},
		{4, 4},
	}

	locations, err := getAsteroidLocations(r)
	assert.NoError(t, err)
	assert.Equal(t, expected, locations)
}

func TestVisibleAsteroidsFrom(t *testing.T) {
	t.Run("example 1", func(t *testing.T) {
		in := `.#..#
		.....
		#####
		....#
		...##`
		in = strings.ReplaceAll(in, "\t", "")
		r := strings.NewReader(in)

		locations, err := getAsteroidLocations(r)
		require.NoError(t, err)
		assert.Equal(t, 7, visibleAsteroidsFrom(point{1, 0}, locations))
	})

	t.Run("example 2", func(t *testing.T) {
		in := `......#.#.
		#..#.#....
		..#######.
		.#.#.###..
		.#..#.....
		..#....#.#
		#..#....#.
		.##.#..###
		##...#..#.
		.#....####`
		in = strings.ReplaceAll(in, "\t", "")
		r := strings.NewReader(in)

		locations, err := getAsteroidLocations(r)
		require.NoError(t, err)
		assert.Equal(t, 33, visibleAsteroidsFrom(point{5, 8}, locations))
	})
}

func TestVaporizeAsteroids(t *testing.T) {
	t.Run("simple", func(t *testing.T) {
		in := `..#.#
		.....
		#####
		....#
		...##`
		in = strings.ReplaceAll(in, "\t", "")
		r := strings.NewReader(in)

		locations, err := getAsteroidLocations(r)
		require.NoError(t, err)

		expected := []point{
			{2, 0},
			{4, 0},
			{3, 2},
			{4, 3},
			{4, 4},
			{3, 4},
			{1, 2},
			{4, 2},
			{0, 2},
		}

		result := vaporizeAsteroids(point{2, 2}, locations)
		assert.Equal(t, expected, result)
	})

	t.Run("a little harder", func(t *testing.T) {
		in := `.#..##.###...#######
		##.############..##.
		.#.######.########.#
		.###.#######.####.#.
		#####.##.#.##.###.##
		..#####..#.#########
		####################
		#.####....###.#.#.##
		##.#################
		#####.##.###..####..
		..######..##.#######
		####.##.####...##..#
		.#####..#.######.###
		##...#.##########...
		#.##########.#######
		.####.#.###.###.#.##
		....##.##.###..#####
		.#.#.###########.###
		#.#.#.#####.####.###
		###.##.####.##.#..##`
		in = strings.ReplaceAll(in, "\t", "")
		r := strings.NewReader(in)
		locations, err := getAsteroidLocations(r)
		require.NoError(t, err)

		asteroid, count := findBestLocation(locations)
		assert.Equal(t, point{11, 13}, asteroid)
		assert.Equal(t, 210, count)

		result := vaporizeAsteroids(asteroid, locations)
		assert.Equal(t, point{8, 2}, result[199])
	})
}
