package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/djlechuck/advent-of-code/utils"
)

type freshRange struct {
	start, end int
}

func main() {
	in := utils.ParseInput("inputs/2025/05.txt")
	fr, ids := processInput(in)

	p1v, p1d := partOne(fr, ids)
	fmt.Printf("Part one: %d (%s)\n", p1v, p1d)
}

func processInput(in utils.Input) ([]freshRange, []int) {
	var fr []freshRange
	var ids []int
	buildRanges := true

	for _, line := range in {
		if "" == line {
			buildRanges = false
			continue
		}

		if buildRanges {
			p := strings.Split(line, "-")
			fr = append(fr, freshRange{start: utils.CastInt(p[0]), end: utils.CastInt(p[1])})
		} else {
			ids = append(ids, utils.CastInt(line))
		}
	}

	return fr, ids
}

func partOne(fr []freshRange, ids []int) (int, string) {
	t := time.Now()
	nbOK := 0

	for _, id := range ids {
		for _, r := range fr {
			if id >= r.start && id <= r.end {
				nbOK++
				break
			}
		}
	}

	return nbOK, time.Since(t).String()
}

func partTwo() (int, string) {
	t := time.Now()

	return 0, time.Since(t).String()
}
