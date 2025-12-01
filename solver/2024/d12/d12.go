package main

import (
	"fmt"
	"slices"
	"time"

	"github.com/djlechuck/advent-of-code/utils"
)

type coord struct{ x, y int }
type plots map[string][]coord

func main() {
}

func processInput(in utils.Input) plots {
	var p = plots{}

	for y, line := range in {
		for x, c := range line {
			p[string(c)] = append(p[string(c)], coord{x, y})
		}
	}

	return p
}

func partOne(p plots) (int, string) {
	t := time.Now()

	for pl, coords := range p {
		var pc []coord
		s := 0
		for _, c := range coords {
			ci := slices.Index(pc, c)
			if ci != -1 {
				s++

				continue
			}

			pc = append(pc, c)
		}
		fmt.Printf("plant %s -> %d\n", pl, s)
	}

	return 0, time.Since(t).String()
}

func partTwo() (int, string) {
	t := time.Now()

	return 0, time.Since(t).String()
}
