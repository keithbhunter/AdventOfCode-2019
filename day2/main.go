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

	instructions[1] = 12
	instructions[2] = 2

	result := executeInstructions(instructions)
	fmt.Printf("%v\n", result[0])
}

func executeInstructions(instructions []int) []int {
	for i := 0; i < len(instructions); i += 4 {
		switch instructions[i] {
		case 1:
			operand1 := instructions[i+1]
			operand2 := instructions[i+2]
			resultPosition := instructions[i+3]
			instructions[resultPosition] = instructions[operand1] + instructions[operand2]
		case 2:
			operand1 := instructions[i+1]
			operand2 := instructions[i+2]
			resultPosition := instructions[i+3]
			instructions[resultPosition] = instructions[operand1] * instructions[operand2]
		case 99:
			break
		}
	}
	return instructions
}
