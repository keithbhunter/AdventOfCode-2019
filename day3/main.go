package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strings"

	"github.com/pkg/errors"
)

func main() {
	b, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(errors.Wrap(err, "could not read file"))
	}

	strings := strings.Split(strings.Trim(string(b), "\n"), "\n")
	if len(strings) != 2 {
		panic(errors.New("too many input strings"))
	}

	w1 := newWire(strings[0])
	w2 := newWire(strings[1])
	s := determineSteps(w1, w2)
	fmt.Printf("steps: %d\n", s)
}

func determineDistance(wire1, wire2 *wire) int {
	d := math.MaxInt32

	for _, coord1 := range wire1.coordinates {
		for _, coord2 := range wire2.coordinates {
			if coord1 == coord2 {
				dist := distance(coord1)
				if dist < d {
					d = dist
				}
			}
		}
	}

	return d
}

func distance(c coordinate) int {
	return int(math.Abs(float64(c.x)) + math.Abs(float64(c.y)))
}

func determineSteps(wire1, wire2 *wire) int {
	minSteps := math.MaxInt32

	for i := 0; i < len(wire1.coordinates); i++ {
		coord1 := wire1.coordinates[i]

		for j := 0; j < len(wire2.coordinates); j++ {
			coord2 := wire2.coordinates[j]

			if coord1 == coord2 {
				steps := i + j + 2
				if steps < minSteps {
					minSteps = steps
				}
			}
		}
	}

	return minSteps
}
