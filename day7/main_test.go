package main

import (
	"io/ioutil"
	"strconv"
	"strings"
	"testing"

	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
)

func TestPart1(t *testing.T) {
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

	highestSignal := 0
	phases := []int{0, 1, 2, 3, 4}
	for _, phaseSeq := range permutations(phases) {
		if output := tryAmplifierPhaseSequence(phaseSeq, instructions); output > highestSignal {
			highestSignal = output
		}
	}

	assert.Equal(t, 255590, highestSignal)
}

// Max thruster signal 43210 (from phase setting sequence 4,3,2,1,0):
// 3,15,3,16,1002,16,10,16,1,16,15,15,4,15,99,0,0
func TestExample1(t *testing.T) {
	s := []int{4, 3, 2, 1, 0}
	input := []int{3, 15, 3, 16, 1002, 16, 10, 16, 1, 16, 15, 15, 4, 15, 99, 0, 0}
	output := tryAmplifierPhaseSequence(s, input)
	assert.Equal(t, 43210, output)
}

// Max thruster signal 54321 (from phase setting sequence 0,1,2,3,4):
// 3,23,3,24,1002,24,10,24,1002,23,-1,23,101,5,23,23,1,24,23,23,4,23,99,0,0
func TestExample2(t *testing.T) {
	s := []int{0, 1, 2, 3, 4}
	input := []int{3, 23, 3, 24, 1002, 24, 10, 24, 1002, 23, -1, 23, 101, 5, 23, 23, 1, 24, 23, 23, 4, 23, 99, 0, 0}
	output := tryAmplifierPhaseSequence(s, input)
	assert.Equal(t, 54321, output)
}

// Max thruster signal 65210 (from phase setting sequence 1,0,4,3,2):
// 3,31,3,32,1002,32,10,32,1001,31,-2,31,1007,31,0,33,1002,33,7,33,1,33,31,31,1,32,31,31,4,31,99,0,0,0
func TestExample3(t *testing.T) {
	s := []int{1, 0, 4, 3, 2}
	input := []int{3, 31, 3, 32, 1002, 32, 10, 32, 1001, 31, -2, 31, 1007, 31, 0, 33, 1002, 33, 7, 33, 1, 33, 31, 31, 1, 32, 31, 31, 4, 31, 99, 0, 0, 0}
	output := tryAmplifierPhaseSequence(s, input)
	assert.Equal(t, 65210, output)
}

func TestPermutations(t *testing.T) {
	arr := []int{0, 1, 2}
	expected := [][]int{
		[]int{0, 1, 2},
		[]int{1, 0, 2},
		[]int{2, 0, 1},
		[]int{0, 2, 1},
		[]int{1, 2, 0},
		[]int{2, 1, 0},
	}
	assert.Equal(t, expected, permutations(arr))
}

func TestSwap(t *testing.T) {
	arr := []int{0, 1, 2, 3, 4, 5}
	swap(arr, 2, 5)
	expected := []int{0, 1, 5, 3, 4, 2}
	assert.Equal(t, expected, arr)
}
