package main

import (
	"io"
	"math"
	"sort"
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

func vaporizeAsteroids(p point, asteroids []point) []point {
	m := map[slopeVector][]point{}

	// build the map of asteroids by slope
	for _, asteroid := range asteroids {
		if asteroid == p {
			continue
		}

		s := slopeBetweenPoints(p, asteroid)
		a, ok := m[s]
		if !ok {
			m[s] = []point{asteroid}
			continue
		}

		new := make([]point, len(a)+1)
		d := distanceBetweenPoints(p, asteroid)
		insert := asteroid

		for i := 0; i < len(new); i++ {
			if i == len(new)-1 {
				new[i] = insert
				break
			}

			d1 := distanceBetweenPoints(p, a[i])
			if d1 > d {
				new[i] = insert
				insert = a[i]
				d = d1
			} else {
				new[i] = a[i]
			}
		}
		m[s] = new
	}

	// sort keys clockwise
	keys := make([]slopeVector, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Sort(bySlope(keys))

	// loop clockwise, moving 1 asteroid at a time from the map
	// into lasered
	lasered := []point{}
	for len(lasered) < len(asteroids)-1 {
		for _, k := range keys {
			as, ok := m[k]
			if !ok {
				continue
			}

			lasered = append(lasered, as[0])
			as = as[1:]
			if len(as) == 0 {
				delete(m, k)
			} else {
				m[k] = as
			}
		}
	}

	return lasered
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
	y = -y
	return slopeVector{slope: y / x, positiveX: x >= 0, positiveY: y >= 0}
}

func distanceBetweenPoints(p1, p2 point) float64 {
	x := float64(p2.x) - float64(p1.x)
	y := float64(p2.y) - float64(p1.y)
	return math.Sqrt(x*x + y*y)
}

type slopeVector struct {
	slope                float64
	positiveX, positiveY bool
}

type point struct {
	x, y int
}

// Sort slope vector by slope
type bySlope []slopeVector

func (s bySlope) Len() int {
	return len(s)
}

func (s bySlope) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s bySlope) Less(i, j int) bool {
	if s[i].positiveX == s[j].positiveX && s[i].positiveY == s[j].positiveY {
		return s[i].slope > s[j].slope
	}

	getWeight := func(v slopeVector) int {
		if v.positiveX && v.positiveY {
			return 0
		}
		if v.positiveX && !v.positiveY {
			return 1
		}
		if !v.positiveX && !v.positiveY {
			return 2
		}
		if !v.positiveX && v.positiveY {
			return 3
		}
		panic("not possible")
	}

	return getWeight(s[i]) < getWeight(s[j])
}
