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

		if left == middle {
			if i < len(pw) - 1 {
				right := rune(pw[i+1])
				if middle != right {
					hasDouble = true
				}
			}
		}
	}

	return hasDouble && neverDecrease
}
