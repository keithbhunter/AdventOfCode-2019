package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

func main() {
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

	for i := 0; i < 100; i++ {
		for j := 0; j < 100; j++ {
			instructions[1] = i
			instructions[2] = j
			if executeInstructions(instructions)[0] == 19690720 {
				fmt.Printf("noun: %v; verb: %v\n", i, j)
			}
		}
	}
}

func executeInstructions(instructions []int) []int {
	c := make([]int, len(instructions))
	copy(c, instructions)

	for i := 0; i < len(c); i += 4 {
		switch c[i] {
		case 1:
			operand1 := c[i+1]
			operand2 := c[i+2]
			resultPosition := c[i+3]
			c[resultPosition] = c[operand1] + c[operand2]
		case 2:
			operand1 := c[i+1]
			operand2 := c[i+2]
			resultPosition := c[i+3]
			c[resultPosition] = c[operand1] * c[operand2]
		case 99:
			break
		}
	}
	return c
}
