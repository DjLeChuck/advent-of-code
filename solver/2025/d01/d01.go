package main

import (
	"fmt"
	"strconv"
	"time"

	"github.com/djlechuck/advent-of-code/utils"
)

type instruction struct {
	op     string
	length int
}

func main() {
	in := utils.ParseInput("inputs/2025/01.txt")
	instructions := processInput(in)

	p1v, p1d := partOne(instructions)
	fmt.Printf("Part one: %d (%s)\n", p1v, p1d)
}

func processInput(in utils.Input) []instruction {
	var instructions []instruction
	for _, line := range in {
		op := string(line[0])
		length, _ := strconv.Atoi(line[1:])
		instructions = append(instructions, instruction{op, length})
	}

	return instructions
}

func partOne(ins []instruction) (int, string) {
	t := time.Now()
	dial := 50
	nbZero := 0

	for _, inst := range ins {
		if inst.op == "L" {
			dial -= inst.length
		} else if inst.op == "R" {
			dial += inst.length
		}
		dial = dial % 100

		if dial < 0 {
			dial += 100
		}

		if dial == 0 {
			nbZero++
		}
	}

	return nbZero, time.Since(t).String()
}

func partTwo(ins []instruction) (int, string) {
	t := time.Now()
	dial := 50
	nbZero := 0

	for _, inst := range ins {
		if inst.op == "L" {
			dial -= inst.length
		} else if inst.op == "R" {
			dial += inst.length
		}

		nbZero += dial / 100

		dial = dial % 100

		if dial < 0 {
			dial += 100
		}

		if dial == 0 {
			nbZero++
		}

		fmt.Printf("%+v -> %d - %d\n", inst, dial, nbZero)
	}

	return 0, time.Since(t).String()
}
