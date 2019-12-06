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

	ExecuteInstructions(1, instructions, func(out int) {
		fmt.Printf("%v\n", out)
	})
}

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
		// fmt.Printf("params: %v\n", param)

		switch param.opCode {
		case 1:
			// fmt.Printf("instr: %v\n", c[i:i+4])
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
		case 2:
			// fmt.Printf("instr: %v\n", c[i:i+4])
			operand1 := 0
			if param.op1Mode == ModeImmediate {
				operand1 = c[i+1]
			} else {
				operand1 = c[c[i+1]]
			}

			operand2 := c[c[i+2]]
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
		case 3:
			// fmt.Printf("instr: %v\n", c[i:i+2])
			if param.op1Mode == ModeImmediate {
				panic(errors.New("unexpected immediate"))
			}
			length = 2
			position := c[i+1]
			c[position] = input
		case 4:
			// fmt.Printf("instr: %v\n", c[i:i+2])
			// if param.op1Mode == ModeImmediate {
			// 	panic(errors.New("unexpected immediate"))
			// }
			length = 2
			position := c[i+1]
			outputCallback(c[position])
		case 99:
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
	opCode  int
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
	p.opCode, err = strconv.Atoi(opCode)
	if err != nil {
		panic(errors.Wrap(err, "cannot convert str "+opCode))
	}

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
