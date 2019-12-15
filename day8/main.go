package main

import (
	"io"
)

type image struct {
	layers []*layer
}

func newImage(s size, r io.Reader) *image {
	i := &image{make([]*layer, 0)}
	for {
		l, err := newLayer(s, r)
		if err == io.EOF {
			break
		}
		i.layers = append(i.layers, l)
	}
	return i
}

type layer struct {
	pixels [][]int
}

func newLayer(s size, r io.Reader) (*layer, error) {
	// Initialize to all zeros.
	rows := make([][]int, s.height)
	for i := range rows {
		rows[i] = make([]int, s.width)
	}
	l := &layer{rows}

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

type size struct {
	width, height int
}
