package main

import (
	"reflect"
	"testing"

	"github.com/djlechuck/advent-of-code/utils"
)

var (
	parsedInput = blocks{
		0, 0, ".", ".", ".", 1, 1, 1, ".", ".", ".",
		2, ".", ".", ".", 3, 3, 3, ".", 4, 4, ".", 5,
		5, 5, 5, ".", 6, 6, 6, 6, ".", 7, 7, 7, ".",
		8, 8, 8, 8, 9, 9,
	}
	parsedInputPartTwo = disk{
		{0, 0, true, 2},
		{1, 0, false, 3},
		{2, 1, true, 3},
		{3, 0, false, 3},
		{4, 2, true, 1},
		{5, 0, false, 3},
		{6, 3, true, 3},
		{7, 0, false, 1},
		{8, 4, true, 2},
		{9, 0, false, 1},
		{10, 5, true, 4},
		{11, 0, false, 1},
		{12, 6, true, 4},
		{13, 0, false, 1},
		{14, 7, true, 3},
		{15, 0, false, 1},
		{16, 8, true, 4},
		{17, 9, true, 2},
	}
)

func TestProcessInput(t *testing.T) {
	in := utils.ParseInput("../../../inputs/2024/09_example.txt")
	e := processInput(in)

	if !reflect.DeepEqual(e, parsedInput) {
		t.Errorf("processInput(): got %v, want %v", e, parsedInput)
	}
}

func TestProcessInputPartTwo(t *testing.T) {
	in := utils.ParseInput("../../../inputs/2024/09_example.txt")
	e := processInputPartTwo(in)

	if !reflect.DeepEqual(e, parsedInputPartTwo) {
		t.Errorf("processInput(): got %v, want %v", e, parsedInputPartTwo)
	}
}

func TestPartOne(t *testing.T) {
	expected := 1928

	result, _ := partOne(parsedInput)
	if result != expected {
		t.Errorf("partOne() = %d, want %d", result, expected)
	}
}

func TestPartTwo(t *testing.T) {
	expected := 2858

	result, _ := partTwo(parsedInputPartTwo)
	if result != expected {
		t.Errorf("partTwo() = %d, want %d", result, expected)
	}
}
