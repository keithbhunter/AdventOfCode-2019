package main

import (
	"fmt"
	"time"
)

const endSignal = -999999999

func main() {}

func findMaxThruster(s, instructions []int) int {
	seqs := permutations(s)
	fmt.Printf("number of sequences: %v\n", len(seqs))

	thrust := 0
	i := 0
	for _, seq := range seqs {
		t := tryAmplifierPhaseSequenceWithFeedback(seq, instructions)
		fmt.Printf("%v thurst: %v\n", i, t)
		if t > thrust {
			thrust = t
		}
		i++
	}
	return thrust
}

func tryAmplifierPhaseSequenceWithFeedback(s, instructions []int) int {
	a1 := newAmp(s[0], instructions)
	a2 := newAmp(s[1], instructions)
	a3 := newAmp(s[2], instructions)
	a4 := newAmp(s[3], instructions)
	a5 := newAmp(s[4], instructions)

	a1.run()
	a2.run()
	a3.run()
	a4.run()
	a5.run()
	firstLoop := true

Loop:
	for {
		select {
		case <-a5.finishedExecuting:
			break Loop

		default:
			if firstLoop {
				a1.input <- 0
				firstLoop = false
			}

			select {
			case i := <-a5.output:
				a1.input <- i
			case i := <-a1.output:
				a2.input <- i
			case i := <-a2.output:
				a3.input <- i
			case i := <-a3.output:
				a4.input <- i
			case i := <-a4.output:
				a5.input <- i
			default:
				time.Sleep(5 * time.Millisecond)
				break
			}
		}
	}

	return a5.lastOutput
}

func permutations(arr []int) [][]int {
	var perm func(int, []int) [][]int

	// Heap's algorithm: https://en.wikipedia.org/wiki/Heap%27s_algorithm#
	perm = func(k int, a []int) [][]int {
		if k == 1 {
			c := make([]int, len(a))
			copy(c, a)
			return [][]int{c}
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

func swap(arr []int, p1, p2 int) {
	tmp := arr[p2]
	arr[p2] = arr[p1]
	arr[p1] = tmp
}
