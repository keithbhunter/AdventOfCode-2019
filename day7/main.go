package main

import (
	intcode "github.com/keithbhunter/AdventOfCode-2019/day5"
	"github.com/pkg/errors"
)

func main() {}

func tryAmplifierPhaseSequence(s, instructions []int) int {
	input := 0
	for _, phase := range s {
		inputCount := 0
		inputCallback := func() int {
			if inputCount == 0 {
				inputCount++
				return phase
			}
			if inputCount == 1 {
				return input
			}
			panic(errors.Errorf("asked for input too many (%d) times", inputCount))
		}
		outputCallback := func(out int) { input = out }
		intcode.ExecuteInstructions(inputCallback, instructions, outputCallback)
	}
	return input
}

func permutations(arr []int) [][]int {
	var perm func(int, []int) [][]int

	// Heap's algorithm: https://en.wikipedia.org/wiki/Heap%27s_algorithm#
	perm = func(k int, a []int) [][]int {
		if k == 1 {
			c := make([]int, len(a))
			copy(c, a)
			return [][]int{c}
		}

		r := perm(k-1, a)

		for i := 0; i < k-1; i++ {
			if k%2 == 0 {
				swap(a, i, k-1)
			} else {
				swap(a, 0, k-1)
			}
			r = append(r, perm(k-1, a)...)
		}
		return r
	}

	return perm(len(arr), arr)
}

func swap(arr []int, p1, p2 int) {
	tmp := arr[p2]
	arr[p2] = arr[p1]
	arr[p1] = tmp
}
