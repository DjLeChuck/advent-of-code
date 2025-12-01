package main

import (
	"fmt"
	"math"
	"regexp"
	"time"

	"github.com/djlechuck/advent-of-code/utils"
)

type clawButton struct {
	deltaX, deltaY int
}
type prize struct {
	x, y int
}
type clawMachine struct {
	btnA  clawButton
	btnB  clawButton
	prize prize
}

func main() {
	in := utils.ParseInput("inputs/2024/13.txt")
	m := processInput(in)

	p1v, p1d := partOne(m)
	fmt.Printf("Part one: %d - elapsed: %s\n", p1v, p1d)
}

func processInput(in utils.Input) []clawMachine {
	var m []clawMachine
	btnRegex := regexp.MustCompile(`Button ([AB]): X\+(\d+), Y\+(\d+)`)
	prizeRegex := regexp.MustCompile(`Prize: X=(\d+), Y=(\d+)`)
	nm := clawMachine{}

	for _, line := range in {
		if btnRegex.MatchString(line) {
			matches := btnRegex.FindStringSubmatch(line)
			if "A" == matches[1] {
				nm.btnA = clawButton{utils.CastInt(matches[2]), utils.CastInt(matches[3])}
			} else if "B" == matches[1] {
				nm.btnB = clawButton{utils.CastInt(matches[2]), utils.CastInt(matches[3])}
			}
		} else if prizeRegex.MatchString(line) {
			matches := prizeRegex.FindStringSubmatch(line)
			nm.prize = prize{utils.CastInt(matches[1]), utils.CastInt(matches[2])}

			m = append(m, nm)
		}
	}

	return m
}

func partOne(m []clawMachine) (int, string) {
	t := time.Now()

	for _, cm := range m {
		maxAX := math.Floor(float64(cm.prize.x / cm.btnA.deltaX))
		maxAY := math.Floor(float64(cm.prize.y / cm.btnA.deltaY))
		maxBX := math.Floor(float64(cm.prize.x / cm.btnB.deltaX))
		maxBY := math.Floor(float64(cm.prize.y / cm.btnB.deltaY))

		fmt.Printf("Max A -> X: %.0f ; Y: %.0f\n", maxAX, maxAY)
		fmt.Printf("Max B -> X: %.0f ; Y: %.0f\n\n", maxBX, maxBY)
	}

	return 0, time.Since(t).String()
}

func partTwo() (int, string) {
	t := time.Now()

	return 0, time.Since(t).String()
}
