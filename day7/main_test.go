package main

import (
	"fmt"
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

	s := []int{9, 8, 7, 6, 5}
	seq := permutations(s)

	output := tryAmplifierPhaseSequenceWithFeedback(seq[0], instructions)
	assert.Equal(t, 139629729, output)
}

// Max thruster signal 139629729 (from phase setting sequence 9,8,7,6,5):
// 3,26,1001,26,-4,26,3,27,1002,27,2,27,1,27,26,27,4,27,1001,28,-1,28,1005,28,6,99,0,0,5
func TestExample4(t *testing.T) {
	s := []int{9, 8, 7, 6, 5}
	input := []int{3, 26, 1001, 26, -4, 26, 3, 27, 1002, 27, 2, 27, 1, 27, 26, 27, 4, 27, 1001, 28, -1, 28, 1005, 28, 6, 99, 0, 0, 5}

	output := findMaxThruster(s, input)
	assert.Equal(t, 139629729, output)
}

// Max thruster signal 18216 (from phase setting sequence 9,7,8,5,6):
// 3,52,1001,52,-5,52,3,53,1,52,56,54,1007,54,5,55,1005,55,26,1001,54,-5,54,1105,1,12,1,53,54,53,1008,54,0,55,1001,55,1,55,2,53,55,53,4,53,1001,56,-1,56,1005,56,6,99,0,0,0,0,10
func TestExample5(t *testing.T) {
	s := []int{9, 7, 8, 5, 6}
	input := []int{3, 52, 1001, 52, -5, 52, 3, 53, 1, 52, 56, 54, 1007, 54, 5, 55, 1005, 55, 26, 1001, 54, -5, 54, 1105, 1, 12, 1, 53, 54, 53, 1008, 54, 0, 55, 1001, 55, 1, 55, 2, 53, 55, 53, 4, 53, 1001, 56, -1, 56, 1005, 56, 6, 99, 0, 0, 0, 0, 10}
	output := tryAmplifierPhaseSequenceWithFeedback(s, input)
	assert.Equal(t, 18216, output)
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
