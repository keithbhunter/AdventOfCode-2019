package main

import (
	"fmt"
	"github.com/pkg/errors"
)

type color int

const (
	black color = 0
	white color = 1
)

type direction int

const (
	turnLeft  direction = 0
	turnRight direction = 1
)

func (d direction) String() string {
	switch d {
	case turnLeft:
		return "turn left"
	case turnRight:
		return "turn right"
	}
	panic(errors.Errorf("unexpected direction %d", d))
}

type facing int

const (
	left facing = iota
	right
	up
	down
)

func (f facing) String() string {
	switch f {
	case left:
		return "left"
	case right:
		return "right"
	case up:
		return "up"
	case down:
		return "down"
	}
	panic(errors.Errorf("unexpected facing %d\n", f))
}

type point struct {
	x, y int
}

func (p point) String() string {
	return fmt.Sprintf("(%d, %d)", p.x, p.y)
}

// type size struct {
// 	width, height int
// }

// type ship struct {
// 	panels [][]color
// }

// func newShip(s size) *ship {
// 	rows := make([][]color, s.height)
// 	for i := range rows {
// 		rows[i] = make([]color, s.width)
// 	}
// 	return &ship{rows}
// }

// func (s *ship) pixel(row, col int) color {
// 	return s.panels[row][col]
// }

// func (s *ship) setPanelColor(c color, row, col int) {
// 	s.panels[row][col] = c
// }

type robot struct {
	computer      *computer
	panels        map[point]color
	facing        facing
	position      point
	outputCounter int
}

func newRobot(instructions []int) *robot {
	c := newComputer(instructions)

	r := &robot{
		computer: c,
		panels:   make(map[point]color),
		facing:   up,
	}

	c.input = func() int {
		var result int
		if color, ok := r.panels[r.position]; ok {
			result = int(color)
		} else {
			result = int(black)
		}
		fmt.Printf("read color at %s: %d\n", r.position, result)
		return result
	}

	c.output = func(o int) {
		r.outputCounter++
		if r.outputCounter%2 == 0 {
			dir := direction(o)
			fmt.Printf("facing %s, %s (%d)\n", r.facing, dir, o)
			r.turn(dir)
			r.moveForward()
			fmt.Printf("now facing %s at position %s\n", r.facing, r.position)
		} else {
			fmt.Printf("setting color at %s to %d\n", r.position, o)
			r.panels[r.position] = color(o)
		}
	}

	return r
}

func (r *robot) run() {
	r.computer.run()
}

func (r *robot) turn(d direction) {
	switch d {
	case turnLeft:
		switch r.facing {
		case left:
			r.facing = down
		case right:
			r.facing = up
		case up:
			r.facing = left
		case down:
			r.facing = right
		}

	case turnRight:
		switch r.facing {
		case left:
			r.facing = up
		case right:
			r.facing = down
		case up:
			r.facing = right
		case down:
			r.facing = left
		}
	}
}

func (r *robot) moveForward() {
	switch r.facing {
	case left:
		r.position.x--
	case right:
		r.position.x++
	case up:
		r.position.y--
	case down:
		r.position.y++
	}
}
