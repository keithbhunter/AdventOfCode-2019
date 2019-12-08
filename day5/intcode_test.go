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

// 3,9,8,9,10,9,4,9,99,-1,8 - Using position mode, consider whether the 
// input is equal to 8; output 1 (if it is) or 0 (if it is not).
func TestExample7(t *testing.T) {
	t.Run("input 1", func(t *testing.T) {
		output := []int{}
		result := ExecuteInstructions(1, []int{3,9,8,9,10,9,4,9,99,-1,8}, func(out int) {
			output = append(output, out)
		})
		assert.Equal(t, []int{3,9,8,9,10,9,4,9,99,0,8}, result)
		assert.Equal(t, []int{0}, output)
	})

	t.Run("input 8", func(t *testing.T) {
		output := []int{}
		result := ExecuteInstructions(8, []int{3,9,8,9,10,9,4,9,99,-1,8}, func(out int) {
			output = append(output, out)
		})
		assert.Equal(t, []int{3,9,8,9,10,9,4,9,99,1,8}, result)
		assert.Equal(t, []int{1}, output)
	})
}

// 3,9,7,9,10,9,4,9,99,-1,8 - Using position mode, consider whether the 
// input is less than 8; output 1 (if it is) or 0 (if it is not).
func TestExample8(t *testing.T) {
	t.Run("input 1", func(t *testing.T) {
		output := []int{}
		result := ExecuteInstructions(1, []int{3,9,7,9,10,9,4,9,99,-1,8}, func(out int) {
			output = append(output, out)
		})
		assert.Equal(t, []int{3,9,7,9,10,9,4,9,99,1,8}, result)
		assert.Equal(t, []int{1}, output)
	})

	t.Run("input 8", func(t *testing.T) {
		output := []int{}
		result := ExecuteInstructions(8, []int{3,9,7,9,10,9,4,9,99,-1,8}, func(out int) {
			output = append(output, out)
		})
		assert.Equal(t, []int{3,9,7,9,10,9,4,9,99,0,8}, result)
		assert.Equal(t, []int{0}, output)
	})
}

// 3,3,1108,-1,8,3,4,3,99 - Using immediate mode, consider whether the 
// input is equal to 8; output 1 (if it is) or 0 (if it is not).
func TestExample9(t *testing.T) {
	t.Run("input 1", func(t *testing.T) {
		output := []int{}
		result := ExecuteInstructions(1, []int{3,3,1108,-1,8,3,4,3,99}, func(out int) {
			output = append(output, out)
		})
		assert.Equal(t, []int{3,3,1108,0,8,3,4,3,99}, result)
		assert.Equal(t, []int{0}, output)
	})

	t.Run("input 8", func(t *testing.T) {
		output := []int{}
		result := ExecuteInstructions(8, []int{3,3,1108,-1,8,3,4,3,99}, func(out int) {
			output = append(output, out)
		})
		assert.Equal(t, []int{3,3,1108,1,8,3,4,3,99}, result)
		assert.Equal(t, []int{1}, output)
	})
}

// 3,3,1107,-1,8,3,4,3,99 - Using immediate mode, consider whether the 
// input is equal to 8; output 1 (if it is) or 0 (if it is not).
func TestExample10(t *testing.T) {
	t.Run("input 1", func(t *testing.T) {
		output := []int{}
		result := ExecuteInstructions(1, []int{3,3,1107,-1,8,3,4,3,99}, func(out int) {
			output = append(output, out)
		})
		assert.Equal(t, []int{3,3,1107,1,8,3,4,3,99}, result)
		assert.Equal(t, []int{1}, output)
	})

	t.Run("input 8", func(t *testing.T) {
		output := []int{}
		result := ExecuteInstructions(8, []int{3,3,1107,-1,8,3,4,3,99}, func(out int) {
			output = append(output, out)
		})
		assert.Equal(t, []int{3,3,1107,0,8,3,4,3,99}, result)
		assert.Equal(t, []int{0}, output)
	})
}

/*
Here are some jump tests that take an input, then output 0 if the input was 
zero or 1 if the input was non-zero:

3,12,6,12,15,1,13,14,13,4,13,99,-1,0,1,9 (using position mode)
3,3,1105,-1,9,1101,0,0,12,4,12,99,1 (using immediate mode)
*/
func TestExample11(t *testing.T) {
	t.Run("instr 6: input 1", func(t *testing.T) {
		output := []int{}
		result := ExecuteInstructions(1, []int{3,12,6,12,15,1,13,14,13,4,13,99,-1,0,1,9}, func(out int) {
			output = append(output, out)
		})
		assert.Equal(t, []int{3,12,6,12,15,1,13,14,13,4,13,99,1,1,1,9}, result)
		assert.Equal(t, []int{1}, output)
	})

	t.Run("instr 6: input 0", func(t *testing.T) {
		output := []int{}
		result := ExecuteInstructions(0, []int{3,12,6,12,15,1,13,14,13,4,13,99,-1,0,1,9}, func(out int) {
			output = append(output, out)
		})
		assert.Equal(t, []int{3,12,6,12,15,1,13,14,13,4,13,99,0,0,1,9}, result)
		assert.Equal(t, []int{0}, output)
	})

	t.Run("instr 5: input 1", func(t *testing.T) {
		output := []int{}
		result := ExecuteInstructions(1, []int{3,3,1105,-1,9,1101,0,0,12,4,12,99,1}, func(out int) {
			output = append(output, out)
		})
		assert.Equal(t, []int{3,3,1105,1,9,1101,0,0,12,4,12,99,1}, result)
		assert.Equal(t, []int{1}, output)
	})

	t.Run("instr 5: input 0", func(t *testing.T) {
		output := []int{}
		result := ExecuteInstructions(0, []int{3,3,1105,-1,9,1101,0,0,12,4,12,99,1}, func(out int) {
			output = append(output, out)
		})
		assert.Equal(t, []int{3,3,1105,0,9,1101,0,0,12,4,12,99,0}, result)
		assert.Equal(t, []int{0}, output)
	})
}

/*
3,21,1008,21,8,20,1005,20,22,107,8,21,20,1006,20,31,1106,0,36,98,0,0,1002,21,
125,20,4,20,1105,1,46,104,999,1105,1,46,1101,1000,1,20,4,20,1105,1,46,98,99

The above example program uses an input instruction to ask for a single number. 
The program will then output 999 if the input value is below 8, output 1000 if 
the input value is equal to 8, or output 1001 if the input value is greater 
than 8.
*/
func TestExample12(t *testing.T) {
	input := []int{3,21,1008,21,8,20,1005,20,22,107,8,21,20,1006,20,31,1106,0,36,98,0,0,1002,21,125,20,4,20,1105,1,46,104,999,1105,1,46,1101,1000,1,20,4,20,1105,1,46,98,99}

	t.Run("input 1", func(t *testing.T) {
		t.SkipNow()
		output := []int{}
		ExecuteInstructions(1, input, func(out int) {
			output = append(output, out)
		})
		assert.Equal(t, []int{999}, output)
	})

	t.Run("input 8", func(t *testing.T) {
		output := []int{}
		ExecuteInstructions(8, input, func(out int) {
			output = append(output, out)
		})
		assert.Equal(t, []int{1000}, output)
	})

	t.Run("input 9", func(t *testing.T) {
		t.SkipNow()
		output := []int{}
		ExecuteInstructions(8, input, func(out int) {
			output = append(output, out)
		})
		assert.Equal(t, []int{1001}, output)
	})
}

func TestNewParam(t *testing.T) {
	p := newParam(1002)
	assert.Equal(t, p.opCode, OpCodeMultiply)
	assert.Equal(t, p.op1Mode, ModePosition)
	assert.Equal(t, p.op2Mode, ModeImmediate)
	assert.Equal(t, p.op3Mode, ModePosition)
}
