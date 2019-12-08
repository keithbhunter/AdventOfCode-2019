package intcode

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

func executeInput() {
	b, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(errors.Wrap(err, "could not read file"))
	}

	strings := strings.Split(strings.Trim(string(b), "\n"), ",")
	instructions := make([]int, 0, len(strings))

	for _, str := range strings {
		instruction, err := strconv.Atoi(str)
		if err != nil {
			panic(errors.Wrapf(err, "could not parse instruction %s", str))
		}
		instructions = append(instructions, instruction)
	}

	ExecuteInstructions(5, instructions, func(out int) {
		fmt.Printf("%v\n", out)
	})
}

type OpCode int

const (
	OpCodeAdd         OpCode = 1
	OpCodeMultiply    OpCode = 2
	OpCodeInput       OpCode = 3
	OpCodeOutput      OpCode = 4
	OpCodeJumpIfTrue  OpCode = 5
	OpCodeJumpIfFalse OpCode = 6
	OpCodeLessThan    OpCode = 7
	OpCodeEquals      OpCode = 8
	OpCodeEnd         OpCode = 99
)

type Mode int

const (
	// ModePosition causes the parameter to be interpreted as a position.
	ModePosition Mode = iota

	// ModeImmediate causes a parameter to be interpreted as a value.
	ModeImmediate
)

func ExecuteInstructions(input int, instructions []int, outputCallback func(int)) []int {
	c := make([]int, len(instructions))
	copy(c, instructions)
	i := 0

Loop:
	for {
		param := newParam(c[i])
		length := 4

		switch param.opCode {
		case OpCodeAdd:
			operand1 := 0
			if param.op1Mode == ModeImmediate {
				operand1 = c[i+1]
			} else {
				operand1 = c[c[i+1]]
			}

			operand2 := 0
			if param.op2Mode == ModeImmediate {
				operand2 = c[i+2]
			} else {
				operand2 = c[c[i+2]]
			}

			resultPosition := c[i+3]
			if param.op3Mode == ModeImmediate {
				panic(errors.New("unexpected immediate"))
			}
			c[resultPosition] = operand1 + operand2

		case OpCodeMultiply:
			operand1 := 0
			if param.op1Mode == ModeImmediate {
				operand1 = c[i+1]
			} else {
				operand1 = c[c[i+1]]
			}

			operand2 := 0
			if param.op2Mode == ModeImmediate {
				operand2 = c[i+2]
			} else {
				operand2 = c[c[i+2]]
			}

			if param.op3Mode == ModeImmediate {
				panic(errors.New("unexpected immediate"))
			}
			resultPosition := c[i+3]
			c[resultPosition] = operand1 * operand2

		case OpCodeInput:
			if param.op1Mode == ModeImmediate {
				panic(errors.New("unexpected immediate"))
			}
			length = 2
			position := c[i+1]
			c[position] = input

		case OpCodeOutput:
			length = 2

			value := 0
			if param.op1Mode == ModeImmediate {
				value = c[i+1]
			} else {
				value = c[c[i+1]]
			}
			outputCallback(value)

		case OpCodeJumpIfTrue:
			operand1 := 0
			if param.op1Mode == ModeImmediate {
				operand1 = c[i+1]
			} else {
				operand1 = c[c[i+1]]
			}

			operand2 := 0
			if param.op2Mode == ModeImmediate {
				operand2 = c[i+2]
			} else {
				operand2 = c[c[i+2]]
			}

			if operand1 != 0 {
				i = operand2
				continue
			}
			length = 3

		case OpCodeJumpIfFalse:
			operand1 := 0
			if param.op1Mode == ModeImmediate {
				operand1 = c[i+1]
			} else {
				operand1 = c[c[i+1]]
			}

			operand2 := 0
			if param.op2Mode == ModeImmediate {
				operand2 = c[i+2]
			} else {
				operand2 = c[c[i+2]]
			}

			if operand1 == 0 {
				i = operand2
				continue
			}
			length = 3

		case OpCodeLessThan:
			operand1 := 0
			if param.op1Mode == ModeImmediate {
				operand1 = c[i+1]
			} else {
				operand1 = c[c[i+1]]
			}

			operand2 := 0
			if param.op2Mode == ModeImmediate {
				operand2 = c[i+2]
			} else {
				operand2 = c[c[i+2]]
			}

			if param.op3Mode == ModeImmediate {
				panic(errors.New("unexpected immediate"))
			}
			resultPosition := c[i+3]

			if operand1 < operand2 {
				c[resultPosition] = 1
			} else {
				c[resultPosition] = 0
			}

		case OpCodeEquals:
			operand1 := 0
			if param.op1Mode == ModeImmediate {
				operand1 = c[i+1]
			} else {
				operand1 = c[c[i+1]]
			}

			operand2 := 0
			if param.op2Mode == ModeImmediate {
				operand2 = c[i+2]
			} else {
				operand2 = c[c[i+2]]
			}

			if param.op3Mode == ModeImmediate {
				panic(errors.New("unexpected immediate"))
			}
			resultPosition := c[i+3]

			if operand1 == operand2 {
				c[resultPosition] = 1
			} else {
				c[resultPosition] = 0
			}

		case OpCodeEnd:
			break Loop

		default:
			panic(errors.Errorf("unexpected param %v, instr %v", param, c[i]))
		}

		if i+length >= len(c) {
			break Loop
		}
		i += length
	}

	return c
}

type parameter struct {
	opCode  OpCode
	op1Mode Mode
	op2Mode Mode
	op3Mode Mode
}

func newParam(input int) parameter {
	str := strconv.Itoa(input)
	len := len(str)

	var err error
	p := parameter{}

	opCode := str
	if len >= 2 {
		opCode = str[len-2:]
	}
	code, err := strconv.Atoi(opCode)
	if err != nil {
		panic(errors.Wrap(err, "cannot convert str "+opCode))
	}
	p.opCode = OpCode(code)

	if len >= 3 {
		op1 := string(str[len-3])
		mode, err := strconv.Atoi(op1)
		p.op1Mode = Mode(mode)
		if err != nil {
			panic(errors.Wrap(err, "cannot convert str "+op1))
		}
	}

	if len >= 4 {
		op2 := string(str[len-4])
		mode, err := strconv.Atoi(op2)
		p.op2Mode = Mode(mode)
		if err != nil {
			panic(errors.Wrap(err, "cannot convert str "+op2))
		}
	}

	if len >= 5 {
		op3 := string(str[len-5])
		mode, err := strconv.Atoi(op3)
		p.op3Mode = Mode(mode)
		if err != nil {
			panic(errors.Wrap(err, "cannot convert str "+op3))
		}
	}

	return p
}
