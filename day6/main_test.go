package main

import (
	"io/ioutil"
	"strings"
	"testing"

	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
)

func TestInput(t *testing.T) {
	b, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(errors.Wrap(err, "could not read file"))
	}

	input := strings.Trim(string(b), "\n")
	om := createOrbitMap(input)
	assert.Equal(t, 142915, om.numberOfOrbits())
	assert.Equal(t, 283, om.findMinimumTransfersTo("YOU", "SAN"))
}

func TestParseOrbitEntry(t *testing.T) {
	c, o := parseOrbitEntry("COM)B")
	assert.Equal(t, "COM", c)
	assert.Equal(t, "B", o)
}

func TestAddOrbit(t *testing.T) {
	om := newOrbitMap()
	om.addOrbit("COM", "B")
	expected := map[string][]string{
		"COM": []string{"B"},
	}
	assert.Equal(t, expected, om.satellites)
}

func TestCreateOrbitMap(t *testing.T) {
	input := `COM)B
	B)C
	C)D
	D)E
	E)F
	B)G
	G)H
	D)I
	E)J
	J)K
	K)L`
	input = strings.ReplaceAll(input, "\t", "")

	om := createOrbitMap(input)
	expected := map[string][]string{
		"COM": []string{"B"},
		"B":   []string{"C", "G"},
		"C":   []string{"D"},
		"D":   []string{"E", "I"},
		"E":   []string{"F", "J"},
		"G":   []string{"H"},
		"J":   []string{"K"},
		"K":   []string{"L"},
	}
	assert.Equal(t, expected, om.satellites)
}

func TestNumberOfOrbiters(t *testing.T) {
	input := `COM)B
	B)C
	C)D
	D)E
	E)F
	B)G
	G)H
	D)I
	E)J
	J)K
	K)L`
	input = strings.ReplaceAll(input, "\t", "")
	om := createOrbitMap(input)
	assert.Equal(t, 2, om.numberOfOrbiters("J"))
}

func TestNumberOfOrbits(t *testing.T) {
	input := `COM)B
	B)C
	C)D
	D)E
	E)F
	B)G
	G)H
	D)I
	E)J
	J)K
	K)L`
	input = strings.ReplaceAll(input, "\t", "")
	om := createOrbitMap(input)
	assert.Equal(t, 42, om.numberOfOrbits())
}

func TestCenterOfOrbiter(t *testing.T) {
	input := `COM)B
	B)C
	C)D
	D)E
	E)F
	B)G
	G)H
	D)I
	E)J
	J)K
	K)L
	K)YOU
	I)SAN`
	input = strings.ReplaceAll(input, "\t", "")
	om := createOrbitMap(input)
	assert.Equal(t, "K", om.centerOfOrbiter("YOU"))
}

func TestOrbitalTransfersTo(t *testing.T) {
	input := `COM)B
	B)C
	C)D
	D)E
	E)F
	B)G
	G)H
	D)I
	E)J
	J)K
	K)L
	K)YOU
	I)SAN`

	t.Run("YOU -> SAN", func(t *testing.T) {
		input = strings.ReplaceAll(input, "\t", "")
		om := createOrbitMap(input)
		assert.Equal(t, []string{}, om.orbitalTransfersTo("YOU", "SAN"))
	})

	t.Run("D -> SAN", func(t *testing.T) {
		input = strings.ReplaceAll(input, "\t", "")
		om := createOrbitMap(input)
		assert.Equal(t, []string{"D", "I", "SAN"}, om.orbitalTransfersTo("D", "SAN"))
	})

	t.Run("B -> SAN", func(t *testing.T) {
		input = strings.ReplaceAll(input, "\t", "")
		om := createOrbitMap(input)
		assert.Equal(t, []string{"B", "C", "D", "I", "SAN"}, om.orbitalTransfersTo("B", "SAN"))
	})
}

func TestFindMinimumTransfersTo(t *testing.T) {
	input := `COM)B
	B)C
	C)D
	D)E
	E)F
	B)G
	G)H
	D)I
	E)J
	J)K
	K)L
	K)YOU
	I)SAN`

	t.Run("YOU -> SAN", func(t *testing.T) {
		input = strings.ReplaceAll(input, "\t", "")
		om := createOrbitMap(input)
		assert.Equal(t, 4, om.findMinimumTransfersTo("YOU", "SAN"))
	})

	t.Run("SAN -> YOU", func(t *testing.T) {
		input = strings.ReplaceAll(input, "\t", "")
		om := createOrbitMap(input)
		assert.Equal(t, 4, om.findMinimumTransfersTo("SAN", "YOU"))
	})

	t.Run("H -> YOU", func(t *testing.T) {
		input = strings.ReplaceAll(input, "\t", "")
		om := createOrbitMap(input)
		assert.Equal(t, 6, om.findMinimumTransfersTo("H", "YOU"))
	})
}
