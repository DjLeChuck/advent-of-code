package main

import (
	"fmt"
	"math"
	"strings"
	"time"

	"github.com/djlechuck/advent-of-code/utils"
)

type coord struct {
	x, y int
}

func main() {
	in := utils.ParseInput("inputs/2025/09.txt")
	c := processInput(in)

	p1v, p1d := partOne(c)
	fmt.Printf("Part one: %d (%s)\n", p1v, p1d)
}

func processInput(in utils.Input) []coord {
	var c []coord

	for _, line := range in {
		parts := strings.Split(line, ",")
		c = append(c, coord{utils.CastInt(parts[0]), utils.CastInt(parts[1])})
	}

	return c
}

func partOne(coords []coord) (int, string) {
	t := time.Now()
	m := 0

	for i, c := range coords {
		for _, c2 := range coords[i+1:] {
			if v := area(c, c2); v > m {
				m = v
			}
		}
	}

	return m, time.Since(t).String()
}

func partTwo() (int, string) {
	t := time.Now()

	return 0, time.Since(t).String()
}

func area(a, b coord) int {
	x := int(math.Abs(float64(a.x-b.x))) + 1
	y := int(math.Abs(float64(a.y-b.y))) + 1

	return x * y
}
