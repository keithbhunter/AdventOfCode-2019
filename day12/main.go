package main

import (
	"regexp"
	"strconv"
)

type point struct {
	x, y, z int
}

type moon struct {
	position point
	velocity point
}

func (m *moon) totalEnergy() int {
	potential := abs(m.position.x) + abs(m.position.y) + abs(m.position.z)
	kinetic := abs(m.velocity.x) + abs(m.velocity.y) + abs(m.velocity.z)
	return potential * kinetic
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func newMoons(in string) []*moon {
	moons := []*moon{}

	exp := regexp.MustCompile("-?[0-9]+")
	numbers := exp.FindAllString(in, -1)

	for i := 0; i < len(numbers); i += 3 {
		x, _ := strconv.Atoi(numbers[i])
		y, _ := strconv.Atoi(numbers[i+1])
		z, _ := strconv.Atoi(numbers[i+2])
		moons = append(moons, &moon{position: point{x, y, z}})
	}

	return moons
}

func simulateMotion(moons []*moon, times int) []*moon {
	for t := 0; t < times; t++ {
		newMoons := []*moon{}
		for _, m := range moons {
			newMoons = append(newMoons,
				&moon{
					position: point{m.position.x, m.position.y, m.position.z},
					velocity: point{m.velocity.x, m.velocity.y, m.velocity.z},
				},
			)
		}

		for i := 0; i < len(moons); i++ {
			for j := i + 1; j < len(moons); j++ {
				one, two := moons[i], moons[j]
				newOne, newTwo := newMoons[i], newMoons[j]

				if one.position.x > two.position.x {
					newOne.velocity.x--
					newTwo.velocity.x++
				} else if one.position.x < two.position.x {
					newOne.velocity.x++
					newTwo.velocity.x--
				}

				if one.position.y > two.position.y {
					newOne.velocity.y--
					newTwo.velocity.y++
				} else if one.position.y < two.position.y {
					newOne.velocity.y++
					newTwo.velocity.y--
				}

				if one.position.z > two.position.z {
					newOne.velocity.z--
					newTwo.velocity.z++
				} else if one.position.z < two.position.z {
					newOne.velocity.z++
					newTwo.velocity.z--
				}
			}
		}

		for _, m := range newMoons {
			m.position.x += m.velocity.x
			m.position.y += m.velocity.y
			m.position.z += m.velocity.z
		}

		moons = newMoons
	}
	return moons
}

func totalEnergyAfter(times int, moons []*moon) int {
	energy := 0
	moons = simulateMotion(moons, times)
	for _, m := range moons {
		energy += m.totalEnergy()
	}
	return energy
}

func permutations(arr []*moon) [][]*moon {
	var perm func(int, []*moon) [][]*moon

	// Heap's algorithm: https://en.wikipedia.org/wiki/Heap%27s_algorithm#
	perm = func(k int, a []*moon) [][]*moon {
		if k == 1 {
			c := make([]*moon, len(a))
			copy(c, a)
			return [][]*moon{c}
		}

		r := perm(k-1, a)

		for i := 0; i < k-1; i++ {
			if k%2 == 0 {
				swap(a, i, k-1)
			} else {
				swap(a, 0, k-1)
			}
			r = append(r, perm(k-1, a)...)
		}
		return r
	}

	return perm(len(arr), arr)
}

func swap(arr []*moon, p1, p2 int) {
	tmp := arr[p2]
	arr[p2] = arr[p1]
	arr[p1] = tmp
}
