package main

import (
	"io/ioutil"
	"strconv"
	"strings"
	"testing"

	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
)

func TestPart2(t *testing.T) {
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

	c := newComputer(instructions)

	c.input = func() int {
		return 2
	}

	out := []int{}
	c.output = func(o int) {
		out = append(out, o)
	}

	c.run()
	assert.Len(t, out, 1)
	assert.Equal(t, 58879, out[0])
}

// 109,1,204,-1,1001,100,1,100,1008,100,16,101,1006,101,0,99 takes no
// input and produces a copy of itself as output.
func TestExample1(t *testing.T) {
	i := []int{109, 1, 204, -1, 1001, 100, 1, 100, 1008, 100, 16, 101, 1006, 101, 0, 99}
	c := newComputer(i)

	c.input = func() int {
		t.Fatal("there should be no input")
		return 0
	}

	out := []int{}
	c.output = func(o int) {
		out = append(out, o)
	}

	c.run()
	assert.Equal(t, i, out)
}

// 1102,34915192,34915192,7,4,7,99,0 should output a 16-digit number.
func TestExample2(t *testing.T) {
	i := []int{1102, 34915192, 34915192, 7, 4, 7, 99, 0}
	c := newComputer(i)

	c.input = func() int {
		t.Fatal("there should be no input")
		return 0
	}

	out := []int{}
	c.output = func(o int) {
		out = append(out, o)
	}

	c.run()
	assert.Equal(t, 1, len(out))
	assert.Equal(t, 1219070632396864, out[0])
}

// 104,1125899906842624,99 should output the large number in the middle.
func TestExample3(t *testing.T) {
	i := []int{104, 1125899906842624, 99}
	c := newComputer(i)

	c.input = func() int {
		t.Fatal("there should be no input")
		return 0
	}

	out := []int{}
	c.output = func(o int) {
		out = append(out, o)
	}

	c.run()
	assert.Equal(t, 1, len(out))
	assert.Equal(t, 1125899906842624, out[0])
}

func TestRelativeInput(t *testing.T) {
	i := []int{109, 1, 203, 2, 1101, 1, 1, 7, 99}
	c := newComputer(i)

	c.input = func() int {
		return 1
	}

	out := []int{}
	c.output = func(o int) {
		out = append(out, o)
	}

	c.run()

	expected := []int{109, 1, 203, 1, 1101, 1, 1, 2, 99}
	assert.Equal(t, expected, c.instructions[0:9])
}
