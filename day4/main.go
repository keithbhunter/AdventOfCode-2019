package main

import "strconv"

func main() {}

func numberOfSuccessfulPasswordsInRange(min, max int) int {
	count := 0
	for i := min; i <= max; i++ {
		if passwordMeetsCriteria(strconv.Itoa(i)) {
			count++
		}
	}
	return count
}

func passwordMeetsCriteria(pw string) bool {
	if len(pw) != 6 {
		return false
	}

	neverDecrease := true
	hasDouble := false

	for i := 1; i < len(pw); i++ {
		left := rune(pw[i-1])
		middle := rune(pw[i])

		if left > middle {
			neverDecrease = false
		}
	}

	runeCount := map[rune]int{}
	for i := 0; i < len(pw); i++ {
		r := rune(pw[i])
		if _, ok := runeCount[r]; ok {
			continue
		}

		count := 1

		for j := i + 1; j < len(pw); j++ {
			r2 := rune(pw[j])
			if r2 == r {
				count++
			} else {
				break
			}
		}

		if count == 2 {
			hasDouble = true
			break
		}
		runeCount[r] = count
	}

	return hasDouble && neverDecrease
}
