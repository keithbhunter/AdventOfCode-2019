package main

import (
	"errors"
	"strings"
)

func main() {}

func createOrbitMap(input string) *orbitMap {
	in := strings.Split(input, "\n")
	m := newOrbitMap()
	for _, i := range in {
		m.addOrbit(parseOrbitEntry(i))
	}
	return m
}

func parseOrbitEntry(s string) (string, string) {
	components := strings.Split(s, ")")
	if len(components) != 2 {
		panic(errors.New("failed to parse orbit string"))
	}
	return components[0], components[1]
}

type orbitMap struct {
	satellites map[string][]string
}

func newOrbitMap() *orbitMap {
	return &orbitMap{map[string][]string{}}
}

func (om *orbitMap) addOrbit(center, orbiter string) {
	satellites, ok := om.satellites[center]
	if !ok {
		satellites = []string{}
	}
	om.satellites[center] = append(satellites, orbiter)
}

func (om *orbitMap) numberOfOrbits() int {
	total := 0
	for center := range om.satellites {
		total += om.numberOfOrbiters(center)
	}
	return total
}

func (om *orbitMap) numberOfOrbiters(center string) int {
	orbiters, ok := om.satellites[center]
	if !ok {
		return 0
	}
	total := len(orbiters)

	for _, orbiter := range orbiters {
		total += om.numberOfOrbiters(orbiter)
	}
	return total
}
