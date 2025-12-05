package main

import (
	"cmp"
	"fmt"
	"slices"
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
	p2v, p2d := partTwo(fr)
	fmt.Printf("Part two: %d (%s)\n", p2v, p2d)
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

func partTwo(fr []freshRange) (int, string) {
	t := time.Now()

	// sort ranges by starting value
	slices.SortFunc(fr, func(a, b freshRange) int {
		return cmp.Compare(a.start, b.start)
	})

	// merge overlapping ranges
	var merged []freshRange
	for _, r := range fr {
		n := len(merged)
		// the first range, or new one starts after the merged ends
		if n == 0 || r.start > merged[n-1].end+1 {
			merged = append(merged, r)
			continue
		}

		// new range overlaps with the merged, update its end
		if r.end > merged[n-1].end {
			merged[n-1].end = r.end
		}
	}

	nbOk := 0
	for _, r := range merged {
		nbOk += r.end - r.start + 1
	}

	return nbOk, time.Since(t).String()
}
