package main

import (
	"math"
	"strconv"

	"github.com/pkg/errors"
)

const log = 5

type amp struct {
	phaseSetting       int
	instructions       []int
	instructionPointer int
	input              chan int
	output             chan int
	lastOutput         int
	finishedExecuting  chan bool
	usedPhaseSetting   bool
}

func newAmp(phaseSetting int, instructions []int) *amp {
	c := make([]int, len(instructions))
	copy(c, instructions)
	return &amp{
		phaseSetting:      phaseSetting,
		instructions:      c,
		input:             make(chan int),
		output:            make(chan int, math.MaxInt32),
		finishedExecuting: make(chan bool, 1),
	}
}

func (a *amp) run() {
	go func() {
		for {
			p, err := newParam(a.instructions[a.instructionPointer])
			if err != nil {
				panic(errors.Wrapf(err, "a%d", a.phaseSetting))
			}
			op1, op2, op3 := a.operandsFromParam(p)

			if p.opCode == opCodeEnd {
				a.finishedExecuting <- true
				return
			}

			signal := a.executeInstruction(p, op1, op2, op3)
			if signal == endSignal {
				return
			}
		}
	}()
}

func (a *amp) operandsFromParam(p parameter) (int, int, int) {
	switch p.opCode {
	case opCodeAdd, opCodeMultiply, opCodeLessThan, opCodeEquals:
		op1 := 0
		if p.op1Mode == modeImmediate {
			op1 = a.instructions[a.instructionPointer+1]
		} else {
			op1 = a.instructions[a.instructions[a.instructionPointer+1]]
		}

		op2 := 0
		if p.op2Mode == modeImmediate {
			op2 = a.instructions[a.instructionPointer+2]
		} else {
			op2 = a.instructions[a.instructions[a.instructionPointer+2]]
		}

		op3 := a.instructions[a.instructionPointer+3]
		return op1, op2, op3

	case opCodeInput:
		return a.instructions[a.instructionPointer+1], 0, 0

	case opCodeOutput:
		op1 := 0
		if p.op1Mode == modeImmediate {
			op1 = a.instructions[a.instructionPointer+1]
		} else {
			op1 = a.instructions[a.instructions[a.instructionPointer+1]]
		}
		return op1, 0, 0

	case opCodeJumpIfTrue, opCodeJumpIfFalse:
		op1 := 0
		if p.op1Mode == modeImmediate {
			op1 = a.instructions[a.instructionPointer+1]
		} else {
			op1 = a.instructions[a.instructions[a.instructionPointer+1]]
		}

		op2 := 0
		if p.op2Mode == modeImmediate {
			op2 = a.instructions[a.instructionPointer+2]
		} else {
			op2 = a.instructions[a.instructions[a.instructionPointer+2]]
		}

		return op1, op2, 0

	case opCodeEnd:
		return 0, 0, 0

	default:
		panic(errors.Errorf("unexpected opcode %v", p.opCode))
	}
}

func (a *amp) executeInstruction(p parameter, op1, op2, op3 int) int {
	switch p.opCode {
	case opCodeAdd:
		a.instructions[op3] = op1 + op2
		a.instructionPointer += p.length

	case opCodeMultiply:
		a.instructions[op3] = op1 * op2
		a.instructionPointer += p.length

	case opCodeInput:
		in := 0
		if !a.usedPhaseSetting {
			in = a.phaseSetting
			a.usedPhaseSetting = true
		} else {
			in = <-a.input
		}

		if in == endSignal {
			return endSignal
		}

		a.instructions[op1] = in
		a.instructionPointer += p.length

	case opCodeOutput:
		a.lastOutput = op1
		a.output <- op1
		a.instructionPointer += p.length

	case opCodeJumpIfTrue:
		if op1 != 0 {
			a.instructionPointer = op2
		} else {
			a.instructionPointer += p.length
		}

	case opCodeJumpIfFalse:
		if op1 == 0 {
			a.instructionPointer = op2
		} else {
			a.instructionPointer += p.length
		}

	case opCodeLessThan:
		if op1 < op2 {
			a.instructions[op3] = 1
		} else {
			a.instructions[op3] = 0
		}
		a.instructionPointer += p.length

	case opCodeEquals:
		if op1 == op2 {
			a.instructions[op3] = 1
		} else {
			a.instructions[op3] = 0
		}
		a.instructionPointer += p.length

	case opCodeEnd:
		return 0

	default:
		panic(errors.Errorf("unexpected param %v", p))
	}

	return 0
}

type opCode int

const (
	opCodeAdd         opCode = 1
	opCodeMultiply    opCode = 2
	opCodeInput       opCode = 3
	opCodeOutput      opCode = 4
	opCodeJumpIfTrue  opCode = 5
	opCodeJumpIfFalse opCode = 6
	opCodeLessThan    opCode = 7
	opCodeEquals      opCode = 8
	opCodeEnd         opCode = 99
)

type mode int

const (
	// modePosition causes the parameter to be interpreted as a position.
	modePosition mode = iota

	// modeImmediate causes a parameter to be interpreted as a value.
	modeImmediate
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
	case opCodeInput, opCodeOutput:
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
