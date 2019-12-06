package intcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInput(t *testing.T) {
	assert.NotPanics(t, func(){ executeInput() })
}

// 1,0,0,0,99 becomes 2,0,0,0,99 (1 + 1 = 2).
func TestExample1(t *testing.T) {
	result := ExecuteInstructions(1, []int{1, 0, 0, 0, 99}, func(output int) {
		t.Fatal("there should be no output")
	})
	assert.Equal(t, []int{2, 0, 0, 0, 99}, result)
}

// 2,3,0,3,99 becomes 2,3,0,6,99 (3 * 2 = 6).
func TestExample2(t *testing.T) {
	result := ExecuteInstructions(1, []int{2, 3, 0, 3, 99}, func(output int) {
		t.Fatal("there should be no output")
	})
	assert.Equal(t, []int{2, 3, 0, 6, 99}, result)
}

// 2,4,4,5,99,0 becomes 2,4,4,5,99,9801 (99 * 99 = 9801).
func TestExample3(t *testing.T) {
	result := ExecuteInstructions(1, []int{2, 4, 4, 5, 99, 0}, func(output int) {
		t.Fatal("there should be no output")
	})
	assert.Equal(t, []int{2, 4, 4, 5, 99, 9801}, result)
}

// 1,1,1,4,99,5,6,0,99 becomes 30,1,1,4,2,5,6,0,99.
func TestExample4(t *testing.T) {
	result := ExecuteInstructions(1, []int{1, 1, 1, 4, 99, 5, 6, 0, 99}, func(output int) {
		t.Fatal("there should be no output")
	})
	assert.Equal(t, []int{30, 1, 1, 4, 2, 5, 6, 0, 99}, result)
}

// 3,0,4,0,99 outputs whatever it gets as input, then halts
func TestExample5(t *testing.T) {
	output := []int{}
	result := ExecuteInstructions(1, []int{3, 0, 4, 0, 99}, func(out int) {
		output = append(output, out)
	})
	assert.Equal(t, []int{1, 0, 4, 0, 99}, result)
	assert.Equal(t, []int{1}, output)
}

// 1002,4,3,4,33
func TestExample6(t *testing.T) {
	result := ExecuteInstructions(1, []int{1002, 4, 3, 4, 33}, func(out int) {
		t.Fatal("there should be no output")
	})
	assert.Equal(t, []int{1002, 4, 3, 4, 99}, result)
}

func TestNewParam(t *testing.T) {
	p := newParam(1002)
	assert.Equal(t, p.opCode, 2)
	assert.Equal(t, p.op1Mode, ModePosition)
	assert.Equal(t, p.op2Mode, ModeImmediate)
	assert.Equal(t, p.op3Mode, ModePosition)
}
