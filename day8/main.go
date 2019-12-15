package main

import (
	"fmt"
	"io"
)

type image struct {
	layers []*layer
	render *layer
}

func newImage(s size, r io.Reader) *image {
	i := &image{
		layers: make([]*layer, 0),
		render: emptyLayer(s),
	}

	for {
		l, err := newLayer(s, r)
		if err == io.EOF {
			break
		}
		i.layers = append(i.layers, l)
	}

	for row := 0; row < s.height; row++ {
		for col := 0; col < s.width; col++ {
			i.decodePixel(row, col)
		}
	}

	return i
}

func (i *image) decodePixel(row, col int) {
	for _, l := range i.layers {
		p := l.pixel(row, col)
		if p == 2 {
			continue
		}
		i.render.setPixel(p, row, col)
		return
	}
}

func (i *image) print() {
	fmt.Print("\n")
	for row := 0; row < i.render.size.height; row++ {
		for col := 0; col < i.render.size.width; col++ {
			p := i.render.pixel(row, col)
			if p == 1 {
				fmt.Print("1")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Print("\n")
	}
	fmt.Print("\n")
}

type layer struct {
	pixels [][]int
	size   size
}

func emptyLayer(s size) *layer {
	rows := make([][]int, s.height)
	for i := range rows {
		rows[i] = make([]int, s.width)
	}
	return &layer{rows, s}
}

func newLayer(s size, r io.Reader) (*layer, error) {
	l := emptyLayer(s)

	// Load all the pixels.
	for i := 0; i < s.height; i++ {
		for j := 0; j < s.width; j++ {
			b := make([]byte, 1)
			_, err := r.Read(b)
			if err == io.EOF {
				return l, err
			}

			// This assumes every byte represents a single ASCII digit.
			l.pixels[i][j] = int(b[0] - '0')
		}
	}
	return l, nil
}

func (l *layer) numberOfDigit(d int) int {
	count := 0
	for _, row := range l.pixels {
		for _, digit := range row {
			if d == digit {
				count++
			}
		}
	}
	return count
}

func (l *layer) pixel(row, col int) int {
	return l.pixels[row][col]
}

func (l *layer) setPixel(p, row, col int) {
	l.pixels[row][col] = p
}

type size struct {
	width, height int
}
