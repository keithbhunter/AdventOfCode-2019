package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIntcode(t *testing.T) {
	t.Run("1,0,0,0,99 becomes 2,0,0,0,99 (1 + 1 = 2)", func(t *testing.T) {
		instr := []int{1, 0, 0, 0, 99}
		expected := []int{2, 0, 0, 0, 99}

		a := newAmp(0, instr)
		a.run()

		assert.Equal(t, expected, a.instructions)
	})

	t.Run("2,3,0,3,99 becomes 2,3,0,6,99 (3 * 2 = 6)", func(t *testing.T) {
		instr := []int{2, 3, 0, 3, 99}
		expected := []int{2, 3, 0, 6, 99}

		a := newAmp(0, instr)
		a.run()

		assert.Equal(t, expected, a.instructions)
	})

	t.Run("2,4,4,5,99,0 becomes 2,4,4,5,99,9801 (99 * 99 = 9801)", func(t *testing.T) {
		instr := []int{2, 4, 4, 5, 99, 0}
		expected := []int{2, 4, 4, 5, 99, 9801}

		a := newAmp(0, instr)
		a.run()

		assert.Equal(t, expected, a.instructions)
	})

	t.Run("1,1,1,4,99,5,6,0,99 becomes 30,1,1,4,2,5,6,0,99", func(t *testing.T) {
		instr := []int{1, 1, 1, 4, 99, 5, 6, 0, 99}
		expected := []int{30, 1, 1, 4, 2, 5, 6, 0, 99}

		a := newAmp(0, instr)
		a.run()

		assert.Equal(t, expected, a.instructions)
	})

	t.Run("3,0,4,0,99 outputs whatever it gets as input, then halts", func(t *testing.T) {
		instr := []int{3, 0, 4, 0, 99}
		expected := []int{1, 0, 4, 0, 99}

		a := newAmp(0, instr)
		a.run()
		a.input <- 1
		out := <-a.output

		assert.Equal(t, expected, a.instructions)
		assert.Equal(t, 1, out)
	})
}

func TestNewParam(t *testing.T) {
	p, err := newParam(1002)
	assert.NoError(t, err)
	assert.Equal(t, p.opCode, opCodeMultiply)
	assert.Equal(t, p.op1Mode, modePosition)
	assert.Equal(t, p.op2Mode, modeImmediate)
	assert.Equal(t, p.op3Mode, modePosition)
}
