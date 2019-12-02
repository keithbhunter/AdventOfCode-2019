package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"

	fcu "github.com/keithbhunter/AdventOfCode-2019/day1/fuelcounterupper"

	"github.com/pkg/errors"
)

func main() {
	b, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(errors.Wrap(err, "could not read file"))
	}

	str := string(b)
	masses := strings.Split(str, "\n")
	sum := 0

	for _, massStr := range masses {
		if massStr == "" {
			continue
		}

		mass, err := strconv.Atoi(massStr)
		if err != nil {
			panic(errors.Wrap(err, "could not convert mass to int"))
		}
		sum += fcu.FuelRequired(mass)
	}

	fmt.Println(sum)
}
