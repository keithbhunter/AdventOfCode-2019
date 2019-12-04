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
	d := determineDistance(w1, w2)
	fmt.Printf("distance: %d\n", d)
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
