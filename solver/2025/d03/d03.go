package main

import (
	"fmt"
	"strconv"
	"time"

	"github.com/djlechuck/advent-of-code/utils"
)

type banks [][]int

func main() {
	in := utils.ParseInput("inputs/2025/03.txt")
	banks := processInput(in)

	p1v, p1d := partOne(banks)
	fmt.Printf("Part one: %d (%s)\n", p1v, p1d)
}

func processInput(in utils.Input) banks {
	banks := make(banks, 0)

	for _, line := range in {
		bank := make([]int, 0)
		for _, char := range line {
			bank = append(bank, int(char-'0'))
		}
		banks = append(banks, bank)
	}

	return banks
}

func partOne(banks banks) (int, string) {
	t := time.Now()

	joltage := 0

	for _, bank := range banks {
		maxI, fJ, sJ := 0, 0, 0

		// find first max
		for i, v := range bank {
			// do not count the last element
			if i == len(bank)-1 {
				continue
			}

			if v > fJ {
				maxI = i
				fJ = v
			}
		}

		// find second max
		for i := maxI + 1; i < len(bank); i++ {
			if bank[i] > sJ {
				sJ = bank[i]
			}
		}

		joltage += utils.CastInt(strconv.Itoa(fJ) + strconv.Itoa(sJ))
	}

	return joltage, time.Since(t).String()
}

func partTwo(banks banks) (int, string) {
	t := time.Now()

	return 0, time.Since(t).String()
}
