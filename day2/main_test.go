package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMain(t *testing.T) {
	main()
}

// 1,0,0,0,99 becomes 2,0,0,0,99 (1 + 1 = 2).
func TestExample1(t *testing.T) {
	result := executeInstructions([]int{1, 0, 0, 0, 99})
	assert.Equal(t, []int{2, 0, 0, 0, 99}, result)
}

// 2,3,0,3,99 becomes 2,3,0,6,99 (3 * 2 = 6).
func TestExample2(t *testing.T) {
	result := executeInstructions([]int{2, 3, 0, 3, 99})
	assert.Equal(t, []int{2, 3, 0, 6, 99}, result)
}

// 2,4,4,5,99,0 becomes 2,4,4,5,99,9801 (99 * 99 = 9801).
func TestExample3(t *testing.T) {
	result := executeInstructions([]int{2, 4, 4, 5, 99, 0})
	assert.Equal(t, []int{2, 4, 4, 5, 99, 9801}, result)
}

// 1,1,1,4,99,5,6,0,99 becomes 30,1,1,4,2,5,6,0,99.
func TestExample4(t *testing.T) {
	result := executeInstructions([]int{1, 1, 1, 4, 99, 5, 6, 0, 99})
	assert.Equal(t, []int{30, 1, 1, 4, 2, 5, 6, 0, 99}, result)
}
