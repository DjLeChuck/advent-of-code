package main

import (
	"reflect"
	"testing"

	"github.com/djlechuck/advent-of-code/utils"
)

var (
	parsedInput = plots{
		"A": {
			{0, 0},
			{1, 0},
			{2, 0},
			{3, 0},
		},
		"B": {
			{0, 1},
			{1, 1},
			{0, 2},
			{1, 2},
		},
		"C": {
			{2, 1},
			{2, 2},
			{3, 2},
			{3, 3},
		},
		"D": {
			{3, 1},
		},
		"E": {
			{0, 3},
			{1, 3},
			{2, 3},
		},
	}
)

func TestProcessInput(t *testing.T) {
	in := utils.ParseInput("../../../inputs/2024/12_example.txt")
	e := processInput(in)

	if !reflect.DeepEqual(e, parsedInput) {
		t.Errorf("processInput(): got %v, want %v", e, parsedInput)
	}
}

func TestPartOne(t *testing.T) {
	expected := 140

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
