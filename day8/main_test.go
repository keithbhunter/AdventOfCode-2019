package main

import (
	"os"
	"math"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestPart1(t *testing.T) {
	f, err := os.Open("input.txt")
	require.NoError(t, err)

	i := newImage(size{25, 6}, f)

	var l *layer
	fewestZeros := math.MaxInt64

	for _, lay := range i.layers {
		numOfZeros := lay.numberOfDigit(0)
		if numOfZeros < fewestZeros {
			fewestZeros = numOfZeros
			l = lay
		}
	}

	result := l.numberOfDigit(1) * l.numberOfDigit(2)
	assert.Equal(t, 1452, result)
}

func TestNewImage(t *testing.T) {
	r := strings.NewReader("123456789012")
	i := newImage(size{3, 2}, r)
	assert.Equal(t, 2, len(i.layers))

	layer1 := [][]int{
		[]int{1, 2, 3},
		[]int{4, 5, 6},
	}
	assert.Equal(t, layer1, i.layers[0].pixels)

	layer2 := [][]int{
		[]int{7, 8, 9},
		[]int{0, 1, 2},
	}
	assert.Equal(t, layer2, i.layers[1].pixels)
}

func TestNumberOfDigit(t *testing.T) {
	r := strings.NewReader("123456789012")
	i := newImage(size{3, 2}, r)

	assert.Equal(t, 0, i.layers[0].numberOfDigit(0))
	assert.Equal(t, 1, i.layers[1].numberOfDigit(0))
}
