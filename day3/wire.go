package main

import (
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

type coordinate struct {
	x, y int
}

type wire struct {
	coordinates []coordinate
}

func newWire(path string) *wire {
	vectors := strings.Split(path, ",")
	w := &wire{make([]coordinate, 0, len(vectors))}
	current := coordinate{0, 0}

	for _, vector := range vectors {
		if len(vector) < 2 {
			panic(errors.New("cannot parse vector " + vector))
		}

		distanceVector := vector[1:]
		distance, err := strconv.Atoi(distanceVector)
		if err != nil {
			panic(errors.Wrapf(err, "cannot convert %s to int", distanceVector))
		}

		var updateCoord func()
		switch vector[0] {
		case 'L':
			updateCoord = func() { current.x-- }
		case 'R':
			updateCoord = func() { current.x++ }
		case 'U':
			updateCoord = func() { current.y++ }
		case 'D':
			updateCoord = func() { current.y-- }
		default:
			panic(errors.New("unknown vector " + vector))
		}

		for i := 0; i < distance; i++ {
			updateCoord()
			w.coordinates = append(w.coordinates, current)
		}
	}

	return w
}
