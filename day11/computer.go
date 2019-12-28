package main

import (
	"math"
	"strconv"

	"github.com/pkg/errors"
)

const log = 5

type computer struct {
	instructions       []int
	instructionPointer int
	input              func() int
	output             func(int)
	relativeBase       int
}

func newComputer(instructions []int) *computer {
	c := make([]int, math.MaxInt16)
	copy(c, instructions)
	return &computer{instructions: c}
}

func (c *computer) run() {
	for {
		p, err := newParam(c.instructions[c.instructionPointer])
		// fmt.Printf("executing %v\n", c.instructions[c.instructionPointer:c.instructionPointer+p.length])
		if err != nil {
			panic(err)
		}
		if p.opCode == opCodeEnd {
			return
		}

		op1, op2, op3 := c.operandsFromParam(p)
		c.executeInstruction(p, op1, op2, op3)
	}
}

func (c *computer) operandsFromParam(p parameter) (int, int, int) {
	getOperand := func(opMode mode, offset int) int {
		if opMode == modeImmediate {
			return c.instructions[c.instructionPointer+offset]
		} else if opMode == modeRelative {
			val := c.instructions[c.instructionPointer+offset]
			pos := c.relativeBase + val

			// When writing to the array, use the relative position as the operand.
			// When reading, the operand is the value at the relative position.
			if p.opCode == opCodeInput || offset == 3 {
				return pos
			}

			return c.instructions[pos]
		} else {
			if offset == 3 {
				return c.instructions[c.instructionPointer+offset]
			}

			pos := c.instructions[c.instructionPointer+offset]
			return c.instructions[pos]
		}
	}

	switch p.opCode {
	case opCodeAdd, opCodeMultiply, opCodeLessThan, opCodeEquals:
		op1 := getOperand(p.op1Mode, 1)
		op2 := getOperand(p.op2Mode, 2)
		op3 := getOperand(p.op3Mode, 3)
		return op1, op2, op3

	case opCodeInput, opCodeOutput, opCodeAdjustRelativeBase:
		op1 := getOperand(p.op1Mode, 1)
		return op1, 0, 0

	case opCodeJumpIfTrue, opCodeJumpIfFalse:
		op1 := getOperand(p.op1Mode, 1)
		op2 := getOperand(p.op2Mode, 2)
		return op1, op2, 0

	case opCodeEnd:
		return 0, 0, 0

	default:
		panic(errors.Errorf("unexpected opcode %v", p.opCode))
	}
}

func (c *computer) executeInstruction(p parameter, op1, op2, op3 int) {
	switch p.opCode {
	case opCodeAdd:
		c.instructions[op3] = op1 + op2

	case opCodeMultiply:
		c.instructions[op3] = op1 * op2

	case opCodeInput:
		c.instructions[op1] = c.input()

	case opCodeOutput:
		c.output(op1)

	case opCodeJumpIfTrue:
		if op1 != 0 {
			c.instructionPointer = op2
			return
		}

	case opCodeJumpIfFalse:
		if op1 == 0 {
			c.instructionPointer = op2
			return
		}

	case opCodeLessThan:
		if op1 < op2 {
			c.instructions[op3] = 1
		} else {
			c.instructions[op3] = 0
		}

	case opCodeEquals:
		if op1 == op2 {
			c.instructions[op3] = 1
		} else {
			c.instructions[op3] = 0
		}

	case opCodeAdjustRelativeBase:
		c.relativeBase += op1

	case opCodeEnd:
		return

	default:
		panic(errors.Errorf("unexpected param %v", p))
	}

	c.instructionPointer += p.length
}

type opCode int

const (
	opCodeAdd                opCode = 1
	opCodeMultiply           opCode = 2
	opCodeInput              opCode = 3
	opCodeOutput             opCode = 4
	opCodeJumpIfTrue         opCode = 5
	opCodeJumpIfFalse        opCode = 6
	opCodeLessThan           opCode = 7
	opCodeEquals             opCode = 8
	opCodeAdjustRelativeBase opCode = 9
	opCodeEnd                opCode = 99
)

type mode int

const (
	modePosition mode = iota
	modeImmediate
	modeRelative
)

type parameter struct {
	opCode  opCode
	op1Mode mode
	op2Mode mode
	op3Mode mode
	length  int
}

func newParam(input int) (parameter, error) {
	str := strconv.Itoa(input)
	len := len(str)

	var err error
	p := parameter{}

	c := str
	if len >= 2 {
		c = str[len-2:]
	}
	code, err := strconv.Atoi(c)
	if err != nil {
		panic(errors.Wrap(err, "cannot convert str "+c))
	}
	p.opCode = opCode(code)

	switch p.opCode {
	case opCodeAdd, opCodeMultiply, opCodeLessThan, opCodeEquals:
		p.length = 4
	case opCodeInput, opCodeOutput, opCodeAdjustRelativeBase:
		p.length = 2
	case opCodeJumpIfTrue, opCodeJumpIfFalse:
		p.length = 3
	case opCodeEnd:
		p.length = 1
	default:
		return parameter{}, errors.Errorf("unexpected opcode %v", code)
	}

	parsemode := func(i int) mode {
		if len >= i {
			op := string(str[len-i])
			m, err := strconv.Atoi(op)
			if err != nil {
				panic(errors.Wrap(err, "cannot convert str "+op))
			}
			return mode(m)
		}
		return modePosition
	}

	p.op1Mode = parsemode(3)
	p.op2Mode = parsemode(4)
	p.op3Mode = parsemode(5)
	return p, nil
}
