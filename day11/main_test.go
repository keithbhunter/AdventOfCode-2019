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

	r := newRobot(instructions)
	r.run()

	assert.Len(t, r.panels, 0)
}
