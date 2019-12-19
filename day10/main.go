package main

import (
	"io"
)

func getAsteroidLocations(r io.Reader) ([]point, error) {
	asteroids := []point{}
	x, y := 0, 0

	for {
		b := make([]byte, 1)
		_, err := r.Read(b)
		if err == io.EOF {
			return asteroids, nil
		}
		if err != nil {
			return nil, err
		}

		r := rune(b[0])
		if r == '#' {
			asteroids = append(asteroids, point{x, y})
		}

		if r == '\n' {
			x = 0
			y++
		} else {
			x++
		}
	}
}

func findBestLocation(asteroids []point) (point, int) {
	v := 0
	p := point{}

	for _, asteroid := range asteroids {
		visible := visibleAsteroidsFrom(asteroid, asteroids)
		if visible > v {
			v = visible
			p = asteroid
		}
	}

	return p, v
}

func visibleAsteroidsFrom(p point, asteroids []point) int {
	m := map[slopeVector]point{}
	for _, asteroid := range asteroids {
		if asteroid == p {
			continue
		}
		m[slopeBetweenPoints(p, asteroid)] = asteroid
	}
	return len(m)
}

func slopeBetweenPoints(p1, p2 point) slopeVector {
	x := float64(p2.x) - float64(p1.x)
	y := float64(p2.y) - float64(p1.y)
	return slopeVector{slope: y / x, positiveX: x >= 0, positiveY: y >= 0}
}

type slopeVector struct {
	slope                float64
	positiveX, positiveY bool
}

type point struct {
	x, y int
}
