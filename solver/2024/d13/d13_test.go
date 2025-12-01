package main

import (
	"reflect"
	"testing"

	"github.com/djlechuck/advent-of-code/utils"
)

var (
	parsedInput = []clawMachine{
		{
			btnA:  clawButton{94, 34},
			btnB:  clawButton{22, 67},
			prize: prize{8400, 5400},
		}, {
			btnA:  clawButton{26, 66},
			btnB:  clawButton{67, 21},
			prize: prize{12748, 12176},
		}, {
			btnA:  clawButton{17, 86},
			btnB:  clawButton{84, 37},
			prize: prize{7870, 6450},
		}, {
			btnA:  clawButton{69, 23},
			btnB:  clawButton{27, 71},
			prize: prize{18641, 10279},
		},
	}
)

func TestProcessInput(t *testing.T) {
	in := utils.ParseInput("../../../inputs/2024/13_example.txt")
	e := processInput(in)

	if !reflect.DeepEqual(e, parsedInput) {
		t.Errorf("processInput(): got %v, want %v", e, parsedInput)
	}
}

func TestPartOne(t *testing.T) {
	expected := 480

	result, _ := partOne(parsedInput)
	if result != expected {
		t.Errorf("partOne() = %d, want %d", result, expected)
	}
}

func TestPartTwo(t *testing.T) {
	expected := 0

	result, _ := partTwo()
	if result != expected {
		t.Errorf("partTwo() = %d, want %d", result, expected)
	}
}
