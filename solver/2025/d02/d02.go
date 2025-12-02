package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/dlclark/regexp2"

	"github.com/djlechuck/advent-of-code/utils"
)

type idsRange struct {
	min, max int
}

func main() {
	in := utils.ParseInput("inputs/2025/02.txt")
	ranges := processInput(in)

	p1v, p1d := partOne(ranges)
	fmt.Printf("Part one: %d (%s)\n", p1v, p1d)
	p2v, p2d := partTwo(ranges)
	fmt.Printf("Part two: %d (%s)\n", p2v, p2d)
}

func processInput(in utils.Input) []idsRange {
	var r = make([]idsRange, 0)

	for _, line := range in {
		for _, sLine := range strings.Split(line, ",") {
			if sLine == "" {
				continue
			}

			parts := strings.Split(sLine, "-")
			minId, _ := strconv.Atoi(parts[0])
			maxId, _ := strconv.Atoi(parts[1])
			r = append(r, idsRange{minId, maxId})
		}
	}

	return r
}

func partOne(ir []idsRange) (int, string) {
	t := time.Now()

	invalids := make(map[int]bool)

	for _, ra := range ir {
		for i := ra.min; i <= ra.max; i++ {
			iStr := strconv.Itoa(i)
			nbChars := len(iStr)

			if nbChars%2 != 0 {
				continue
			}

			l := iStr[0 : nbChars/2]
			r := iStr[nbChars/2:]
			if l == r {
				invalids[i] = true
			}
		}
	}

	nb := 0
	for k := range invalids {
		nb += k
	}

	return nb, time.Since(t).String()
}

func partTwo(ir []idsRange) (int, string) {
	t := time.Now()

	invalids := make(map[int]bool)
	re := regexp2.MustCompile(`^(\d+)\1+$`, 0)

	for _, ra := range ir {
		for i := ra.min; i <= ra.max; i++ {
			iStr := strconv.Itoa(i)

			if isMatch, _ := re.MatchString(iStr); isMatch {
				invalids[i] = true
			}
		}
	}

	nb := 0
	for k := range invalids {
		nb += k
	}

	return nb, time.Since(t).String()
}
