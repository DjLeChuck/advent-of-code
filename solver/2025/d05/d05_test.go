package main

import (
	"reflect"
	"testing"

	"github.com/djlechuck/advent-of-code/utils"
)

var (
	parsedInputRanges = []freshRange{
		{3, 5},
		{10, 14},
		{16, 20},
		{12, 18},
	}
	parsedInputIds = []int{1, 5, 8, 11, 17, 32}
)

func TestProcessInput(t *testing.T) {
	in := utils.ParseInput("../../../inputs/2025/05_example.txt")
	fr, ids := processInput(in)

	if !reflect.DeepEqual(fr, parsedInputRanges) {
		t.Errorf("processInput(): got %v, want %v", fr, parsedInputRanges)
	}

	if !reflect.DeepEqual(ids, parsedInputIds) {
		t.Errorf("processInput(): got %v, want %v", ids, parsedInputIds)
	}
}

func TestPartOne(t *testing.T) {
	expected := 3

	result, _ := partOne(parsedInputRanges, parsedInputIds)
	if result != expected {
		t.Errorf("partOne() = %d, want %d", result, expected)
	}
}

func TestPartTwo(t *testing.T) {
	expected := 14

	result, _ := partTwo(parsedInputRanges)
	if result != expected {
		t.Errorf("partTwo() = %d, want %d", result, expected)
	}
}
